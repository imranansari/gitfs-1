package filesystem

import (
	"os"
	"sync"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/indeedeng/gitfs/pkg/tree"
	rlog "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

func NewDirectory(uid, gid uint32, tree *gitfstree.TreeNode) fs.Node {
	cache := make(map[string]fs.Node)

	return &Directory{
		uid:       uid,
		gid:       gid,
		tree:      tree,
		cacheLock: sync.Mutex{},
		cache:     cache,
	}
}

type Directory struct {
	uid  uint32
	gid  uint32
	tree *gitfstree.TreeNode

	// use a cache so that you always get the same reference back
	// todo: prune cache
	cacheLock sync.Mutex
	cache     map[string]fs.Node
}

func (d *Directory) Lookup(ctx context.Context, name string) (fs.Node, error) {
	d.cacheLock.Lock()
	defer d.cacheLock.Unlock()

	node, ok := d.tree.Child(name)
	if !ok {
		return nil, fuse.ENOENT
	}

	if entry, ok := d.cache[name]; ok {
		return entry, nil
	}

	var directory fs.Node

	if len(node.URL) > 0 {
		myfs := &SynchronizedFilesystem{
			Delegate: memfs.New(),
		}

		gitfs, err := myfs.Chroot(git.GitDirName)
		if err != nil {
			rlog.Errorf("failed to create .git dir")
			return nil, fuse.ENOENT
		}

		storage := filesystem.NewStorage(gitfs, cache.NewObjectLRUDefault())

		rlog.Infof("cloning repository: %s", node.URL)

		// shallow clone for now since we only support read only
		_, err = git.Clone(storage, myfs, &git.CloneOptions{
			URL:   node.URL,
			Depth: 1,
		})

		if err != nil {
			rlog.Errorf("failed to clone repository: %v", err)
			return nil, fuse.ENOENT
		}

		directory = &BillyNode{
			repourl: node.URL,
			fs:      myfs,
			path:    "",
			target:  "",
			user: BillyUser{
				uid: d.uid,
				gid: d.gid,
			},
			mode:  os.ModeDir | defaultPerms,
			size:  0,
			data:  nil,
			mu:    &sync.Mutex{},
			cache: make(map[string]*BillyNode),
		}
	} else {
		directory = NewDirectory(d.uid, d.gid, node)
	}

	d.cache[name] = directory
	return directory, nil
}

func (d *Directory) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	// tree is immutable, no need to lock
	children := d.tree.Children()

	dirents := make([]fuse.Dirent, len(children))

	for i, child := range children {
		dirents[i] = fuse.Dirent{
			Name: child,
		}
	}

	return dirents, nil
}

func (d *Directory) Attr(ctx context.Context, attr *fuse.Attr) error {
	attr.Mode = os.ModeDir | 0755
	return nil
}