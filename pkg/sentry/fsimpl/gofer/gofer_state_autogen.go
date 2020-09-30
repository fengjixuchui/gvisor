// automatically generated by stateify.

package gofer

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *dentryList) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentryList"
}

func (x *dentryList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *dentryList) beforeSave() {}

func (x *dentryList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *dentryList) afterLoad() {}

func (x *dentryList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *dentryEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentryEntry"
}

func (x *dentryEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *dentryEntry) beforeSave() {}

func (x *dentryEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *dentryEntry) afterLoad() {}

func (x *dentryEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func (x *directoryFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.directoryFD"
}

func (x *directoryFD) StateFields() []string {
	return []string{
		"fileDescription",
		"DirectoryFileDescriptionDefaultImpl",
		"off",
		"dirents",
	}
}

func (x *directoryFD) beforeSave() {}

func (x *directoryFD) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fileDescription)
	m.Save(1, &x.DirectoryFileDescriptionDefaultImpl)
	m.Save(2, &x.off)
	m.Save(3, &x.dirents)
}

func (x *directoryFD) afterLoad() {}

func (x *directoryFD) StateLoad(m state.Source) {
	m.Load(0, &x.fileDescription)
	m.Load(1, &x.DirectoryFileDescriptionDefaultImpl)
	m.Load(2, &x.off)
	m.Load(3, &x.dirents)
}

func (x *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.FilesystemType"
}

func (x *FilesystemType) StateFields() []string {
	return []string{}
}

func (x *FilesystemType) beforeSave() {}

func (x *FilesystemType) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *FilesystemType) afterLoad() {}

func (x *FilesystemType) StateLoad(m state.Source) {
}

func (x *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.filesystem"
}

func (x *filesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"mfp",
		"opts",
		"iopts",
		"clock",
		"devMinor",
		"cachedDentries",
		"cachedDentriesLen",
		"syncableDentries",
		"specialFileFDs",
		"syntheticSeq",
	}
}

func (x *filesystem) beforeSave() {}

func (x *filesystem) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsfs)
	m.Save(1, &x.mfp)
	m.Save(2, &x.opts)
	m.Save(3, &x.iopts)
	m.Save(4, &x.clock)
	m.Save(5, &x.devMinor)
	m.Save(6, &x.cachedDentries)
	m.Save(7, &x.cachedDentriesLen)
	m.Save(8, &x.syncableDentries)
	m.Save(9, &x.specialFileFDs)
	m.Save(10, &x.syntheticSeq)
}

func (x *filesystem) afterLoad() {}

func (x *filesystem) StateLoad(m state.Source) {
	m.Load(0, &x.vfsfs)
	m.Load(1, &x.mfp)
	m.Load(2, &x.opts)
	m.Load(3, &x.iopts)
	m.Load(4, &x.clock)
	m.Load(5, &x.devMinor)
	m.Load(6, &x.cachedDentries)
	m.Load(7, &x.cachedDentriesLen)
	m.Load(8, &x.syncableDentries)
	m.Load(9, &x.specialFileFDs)
	m.Load(10, &x.syntheticSeq)
}

func (x *inodeNumber) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.inodeNumber"
}

func (x *inodeNumber) StateFields() []string {
	return nil
}

func (x *filesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.filesystemOptions"
}

func (x *filesystemOptions) StateFields() []string {
	return []string{
		"fd",
		"aname",
		"interop",
		"dfltuid",
		"dfltgid",
		"msize",
		"version",
		"maxCachedDentries",
		"forcePageCache",
		"limitHostFDTranslation",
		"overlayfsStaleRead",
		"regularFilesUseSpecialFileFD",
	}
}

func (x *filesystemOptions) beforeSave() {}

func (x *filesystemOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fd)
	m.Save(1, &x.aname)
	m.Save(2, &x.interop)
	m.Save(3, &x.dfltuid)
	m.Save(4, &x.dfltgid)
	m.Save(5, &x.msize)
	m.Save(6, &x.version)
	m.Save(7, &x.maxCachedDentries)
	m.Save(8, &x.forcePageCache)
	m.Save(9, &x.limitHostFDTranslation)
	m.Save(10, &x.overlayfsStaleRead)
	m.Save(11, &x.regularFilesUseSpecialFileFD)
}

func (x *filesystemOptions) afterLoad() {}

func (x *filesystemOptions) StateLoad(m state.Source) {
	m.Load(0, &x.fd)
	m.Load(1, &x.aname)
	m.Load(2, &x.interop)
	m.Load(3, &x.dfltuid)
	m.Load(4, &x.dfltgid)
	m.Load(5, &x.msize)
	m.Load(6, &x.version)
	m.Load(7, &x.maxCachedDentries)
	m.Load(8, &x.forcePageCache)
	m.Load(9, &x.limitHostFDTranslation)
	m.Load(10, &x.overlayfsStaleRead)
	m.Load(11, &x.regularFilesUseSpecialFileFD)
}

func (x *InteropMode) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.InteropMode"
}

func (x *InteropMode) StateFields() []string {
	return nil
}

func (x *InternalFilesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.InternalFilesystemOptions"
}

func (x *InternalFilesystemOptions) StateFields() []string {
	return []string{
		"LeakConnection",
		"OpenSocketsByConnecting",
	}
}

func (x *InternalFilesystemOptions) beforeSave() {}

func (x *InternalFilesystemOptions) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.LeakConnection)
	m.Save(1, &x.OpenSocketsByConnecting)
}

func (x *InternalFilesystemOptions) afterLoad() {}

func (x *InternalFilesystemOptions) StateLoad(m state.Source) {
	m.Load(0, &x.LeakConnection)
	m.Load(1, &x.OpenSocketsByConnecting)
}

func (x *dentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentry"
}

func (x *dentry) StateFields() []string {
	return []string{
		"vfsd",
		"refs",
		"fs",
		"parent",
		"name",
		"deleted",
		"cached",
		"dentryEntry",
		"children",
		"syntheticChildren",
		"dirents",
		"ino",
		"mode",
		"uid",
		"gid",
		"blockSize",
		"atime",
		"mtime",
		"ctime",
		"btime",
		"size",
		"atimeDirty",
		"mtimeDirty",
		"nlink",
		"mappings",
		"hostFD",
		"cache",
		"dirty",
		"pf",
		"haveTarget",
		"target",
		"endpoint",
		"pipe",
		"locks",
		"watches",
	}
}

func (x *dentry) beforeSave() {}

func (x *dentry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsd)
	m.Save(1, &x.refs)
	m.Save(2, &x.fs)
	m.Save(3, &x.parent)
	m.Save(4, &x.name)
	m.Save(5, &x.deleted)
	m.Save(6, &x.cached)
	m.Save(7, &x.dentryEntry)
	m.Save(8, &x.children)
	m.Save(9, &x.syntheticChildren)
	m.Save(10, &x.dirents)
	m.Save(11, &x.ino)
	m.Save(12, &x.mode)
	m.Save(13, &x.uid)
	m.Save(14, &x.gid)
	m.Save(15, &x.blockSize)
	m.Save(16, &x.atime)
	m.Save(17, &x.mtime)
	m.Save(18, &x.ctime)
	m.Save(19, &x.btime)
	m.Save(20, &x.size)
	m.Save(21, &x.atimeDirty)
	m.Save(22, &x.mtimeDirty)
	m.Save(23, &x.nlink)
	m.Save(24, &x.mappings)
	m.Save(25, &x.hostFD)
	m.Save(26, &x.cache)
	m.Save(27, &x.dirty)
	m.Save(28, &x.pf)
	m.Save(29, &x.haveTarget)
	m.Save(30, &x.target)
	m.Save(31, &x.endpoint)
	m.Save(32, &x.pipe)
	m.Save(33, &x.locks)
	m.Save(34, &x.watches)
}

func (x *dentry) afterLoad() {}

func (x *dentry) StateLoad(m state.Source) {
	m.Load(0, &x.vfsd)
	m.Load(1, &x.refs)
	m.Load(2, &x.fs)
	m.Load(3, &x.parent)
	m.Load(4, &x.name)
	m.Load(5, &x.deleted)
	m.Load(6, &x.cached)
	m.Load(7, &x.dentryEntry)
	m.Load(8, &x.children)
	m.Load(9, &x.syntheticChildren)
	m.Load(10, &x.dirents)
	m.Load(11, &x.ino)
	m.Load(12, &x.mode)
	m.Load(13, &x.uid)
	m.Load(14, &x.gid)
	m.Load(15, &x.blockSize)
	m.Load(16, &x.atime)
	m.Load(17, &x.mtime)
	m.Load(18, &x.ctime)
	m.Load(19, &x.btime)
	m.Load(20, &x.size)
	m.Load(21, &x.atimeDirty)
	m.Load(22, &x.mtimeDirty)
	m.Load(23, &x.nlink)
	m.Load(24, &x.mappings)
	m.Load(25, &x.hostFD)
	m.Load(26, &x.cache)
	m.Load(27, &x.dirty)
	m.Load(28, &x.pf)
	m.Load(29, &x.haveTarget)
	m.Load(30, &x.target)
	m.Load(31, &x.endpoint)
	m.Load(32, &x.pipe)
	m.Load(33, &x.locks)
	m.Load(34, &x.watches)
}

func (x *fileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.fileDescription"
}

func (x *fileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
	}
}

func (x *fileDescription) beforeSave() {}

func (x *fileDescription) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.vfsfd)
	m.Save(1, &x.FileDescriptionDefaultImpl)
	m.Save(2, &x.LockFD)
}

func (x *fileDescription) afterLoad() {}

func (x *fileDescription) StateLoad(m state.Source) {
	m.Load(0, &x.vfsfd)
	m.Load(1, &x.FileDescriptionDefaultImpl)
	m.Load(2, &x.LockFD)
}

func (x *regularFileFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.regularFileFD"
}

func (x *regularFileFD) StateFields() []string {
	return []string{
		"fileDescription",
		"off",
	}
}

func (x *regularFileFD) beforeSave() {}

func (x *regularFileFD) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fileDescription)
	m.Save(1, &x.off)
}

func (x *regularFileFD) afterLoad() {}

func (x *regularFileFD) StateLoad(m state.Source) {
	m.Load(0, &x.fileDescription)
	m.Load(1, &x.off)
}

func (x *dentryPlatformFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.dentryPlatformFile"
}

func (x *dentryPlatformFile) StateFields() []string {
	return []string{
		"dentry",
		"fdRefs",
		"hostFileMapper",
	}
}

func (x *dentryPlatformFile) beforeSave() {}

func (x *dentryPlatformFile) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dentry)
	m.Save(1, &x.fdRefs)
	m.Save(2, &x.hostFileMapper)
}

func (x *dentryPlatformFile) afterLoad() {}

func (x *dentryPlatformFile) StateLoad(m state.Source) {
	m.Load(0, &x.dentry)
	m.Load(1, &x.fdRefs)
	m.Load(2, &x.hostFileMapper)
}

func (x *endpoint) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.endpoint"
}

func (x *endpoint) StateFields() []string {
	return []string{
		"dentry",
		"path",
	}
}

func (x *endpoint) beforeSave() {}

func (x *endpoint) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dentry)
	m.Save(1, &x.path)
}

func (x *endpoint) afterLoad() {}

func (x *endpoint) StateLoad(m state.Source) {
	m.Load(0, &x.dentry)
	m.Load(1, &x.path)
}

func (x *specialFileFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/gofer.specialFileFD"
}

func (x *specialFileFD) StateFields() []string {
	return []string{
		"fileDescription",
		"isRegularFile",
		"seekable",
		"haveQueue",
		"queue",
		"off",
	}
}

func (x *specialFileFD) beforeSave() {}

func (x *specialFileFD) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fileDescription)
	m.Save(1, &x.isRegularFile)
	m.Save(2, &x.seekable)
	m.Save(3, &x.haveQueue)
	m.Save(4, &x.queue)
	m.Save(5, &x.off)
}

func (x *specialFileFD) afterLoad() {}

func (x *specialFileFD) StateLoad(m state.Source) {
	m.Load(0, &x.fileDescription)
	m.Load(1, &x.isRegularFile)
	m.Load(2, &x.seekable)
	m.Load(3, &x.haveQueue)
	m.Load(4, &x.queue)
	m.Load(5, &x.off)
}

func init() {
	state.Register((*dentryList)(nil))
	state.Register((*dentryEntry)(nil))
	state.Register((*directoryFD)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*inodeNumber)(nil))
	state.Register((*filesystemOptions)(nil))
	state.Register((*InteropMode)(nil))
	state.Register((*InternalFilesystemOptions)(nil))
	state.Register((*dentry)(nil))
	state.Register((*fileDescription)(nil))
	state.Register((*regularFileFD)(nil))
	state.Register((*dentryPlatformFile)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*specialFileFD)(nil))
}
