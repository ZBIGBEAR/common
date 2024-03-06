package permission

type PermissionType int64

const (
	NoPerm   PermissionType = iota
	ReadPerm PermissionType = 1 << (iota - 1)
	WritePerm
	ExportPerm
	DeletePerm
	CreatePerm
)

type Perm interface {
	AddPerm(newPerm PermissionType)
	RemovePerm(newPerm PermissionType)
	HasPerm(newPerm PermissionType) bool
}

type perm struct {
	p PermissionType
}

func NewPerm() Perm {
	return &perm{p: NoPerm}
}

func (p *perm) AddPerm(newPerm PermissionType) {
	p.p = p.p | newPerm
}

func (p *perm) RemovePerm(newPerm PermissionType) {
	p.p = p.p & (^newPerm)
}

func (p *perm) HasPerm(newPerm PermissionType) bool {
	return p.p&newPerm == 1
}
