package dto

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/14
 * @Version 1.0.0
 */

type RolePermissionTreeDto struct {
	HaveCheckedKeys    []string             `json:"haveCheckedKeys"`
	RolePermissionTree []RolePermissionTree `json:"rolePermissionTree"`
	PerissionIds       []string             `json:"perissionIds"`
}
type RolePermissionTree struct {
	Type      string               `json:"type"`
	ParentID  string               `json:"parentId"`
	Name      string               `json:"name"`
	Rank      int                  `json:"rank"`
	IsGranted bool                 `json:"isGranted"`
	Children  []RolePermissionTree `json:"children"`
	ID        string               `json:"id"`
	IsBranch  bool                 `json:"-"`
}
