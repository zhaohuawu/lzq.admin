/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : hsf_basic_dev

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 19/02/2022 16:30:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for au_menu
-- ----------------------------
DROP TABLE IF EXISTS `au_menu`;
CREATE TABLE `au_menu`  (
  `Id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ExtraProperties` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `ConcurrencyStamp` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `CreationTime` datetime(6) NOT NULL,
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `LastModificationTime` datetime(6) NULL DEFAULT NULL,
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NOT NULL DEFAULT 0,
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeletionTime` datetime(6) NULL DEFAULT NULL,
  `Code` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Rank` int(11) NOT NULL,
  `Icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ParentId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsBranch` tinyint(1) NOT NULL,
  `Path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Policy` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `IsHidden` tinyint(1) NOT NULL,
  `ModuleId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of au_menu
-- ----------------------------
INSERT INTO `au_menu` VALUES ('075ccc2d-35af-0455-2e9c-39f50af7b4ad', '{}', '6092df602127449f8d298f7cf4812404', '2020-05-09 17:36:16.728482', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-09 18:23:37.651060', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-12 11:26:11.410298', 'TenantManagement', '租户管理', 2, 'peoples', NULL, 1, '/tenantmanagement', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', '{}', '1613b6da5f774156b08d709500614f85', '2020-05-10 19:56:19.113957', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, 'RoleList', '角色管理', 2, NULL, 'cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', 0, 'role-list', '/views/authorization/role-list', 'Authorization.RoleList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('216ad756-bdcc-abd3-c516-39f4ff08b978', '{}', 'ffc3ecdf9e6c44e9b8c55289fa78a460', '2020-05-07 09:59:25.305045', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-09 10:00:05.695690', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'PermissionList', '操作权限', 1, '', 'cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', 0, 'permission-list', '/views/authorize/permission-list', 'Authorization.PermissionList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('3081e84b-656b-cbfd-5c8c-39f78a0fe83d', '{}', '9b22ec2aac124158b576a858788de36e', '2020-09-10 20:00:03.898518', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-21 20:44:41.485719', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'ProcessSettings', '流程设置', 1, NULL, 'f7352ae7-6875-23e3-afe8-39f789a5d9f5', 0, '/workflow/approvaltype-list', '/workflow/approvaltype-list', 'Workflow.ProcessSettings', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('332f6c3f-40cc-2908-99ef-39f4e6b0dfec', '{}', '450ef3e0fc2a431792c4f8b8ddae3f8c', '2020-05-02 16:32:34.797764', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-09 19:11:39.008230', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'MenuList', '菜单管理', 0, '', 'cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', 0, 'menu-list', '/view/authorize/menu-list', 'Authorization.MenuList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('33c763ff-62b2-a190-a3ec-39f50b4e28cf', '{}', '36cb15fe1bd442388517ec2e3700b8c9', '2020-05-09 19:10:42.460452', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-10 14:29:25.960123', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Dashboard', '首页容器', 0, 'home-fill', NULL, 1, '/', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('344f5ead-8a53-ea44-c25e-39f55c3eb97d', '{}', '334a02f02033400ca2ce3050f38c4fe8', '2020-05-25 12:23:05.967979', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-06-03 16:18:55.038789', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-06-28 21:51:57.597826', 'Depts', '部门管理1', 2, NULL, 'cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', 0, '/dept/index', '/views/dept/index', 'Authorization.Depts', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('5c01d798-9d1f-f088-9659-39f50b43e4a3', '{}', 'b09bf397a8ae42b0a3ca3046e3af94f5', '2020-05-09 18:59:29.860932', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-23 16:38:59.057788', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Home', '首页', 0, '', '33c763ff-62b2-a190-a3ec-39f50b4e28cf', 0, 'dashboard', '/views/dashboard/index', 'Dashboard.HomePage', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('753ce348-567c-4e19-22f3-39f534336219', '{}', '70047beb29ab4358bec1c4b354900163', '2020-05-17 17:45:54.171447', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, 'SysUserList', '用户管理', 1, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/sysuser/sysuser-list', '/views/sysuser/sysuser-list', 'Infrastructure.SysUserList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('8c7e1739-c29c-edcb-b541-39f751de902a', '{}', 'a473cc5deed24f619e5af6b06811e980', '2020-08-30 22:07:25.853392', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-03 22:48:24.944178', 'Workflow', '工作流', 3, 'tree', NULL, 1, 'workflow', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('93352161-b3fb-f99e-5cd9-39f55c390fec', '{}', 'eb0df579ceb84e66be9e40e05ac9afaf', '2020-05-25 12:16:54.861810', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, 'Settings', '配置设置', 3, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/settings', '/views/settings/index', 'Infrastructure.Settings', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('a6723edb-74e4-3053-c16b-39f50a959be0', '{}', 'f63b12e3a9304fc19c3863c13576685f', '2020-05-09 15:49:07.818682', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-10 18:15:04.200070', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-10 18:48:01.005610', 'Logs', '日志k', 1, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, 'logs', '/view/logs/index', 'Infrastructure.Logs', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('afa5524e-b236-6607-1911-39f543786863', '{}', 'ce0e14169f124ec29b0bd14517b6fcdb', '2020-05-20 16:55:36.037295', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-27 15:15:50.422349', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Log', '日志', 2, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/log/index', '/view/log/index', 'Infrastructure.Log', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('b5047074-9a0c-8721-b2a0-39f56645d7b0', '{}', '3c955a4511574eb4b6989c4c4ebeed7a', '2020-05-27 11:07:04.710383', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-27 11:28:31.773987', 'Areas', '省市区管理', 4, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/areas/index', '/views/areas/index', 'Infrastructure.Areas', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', '{}', 'bdeda6e5e9784732999556c5425b5f56', '2020-05-05 11:37:18.505680', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-12 21:38:12.978599', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Authorization', '权限管理', 1, 'lock', NULL, 1, '/authorization', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('ce6c3946-3da7-493a-a5c2-1d5ad4444aee', NULL, NULL, '2022-02-12 21:04:05.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, 'ceshi', 'ceshi', 1, '', 'f7352ae7-6875-23e3-afe8-39f789a5d9f5', 0, '/', '/', 'Workflow.ceshi', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('d4ff0631-df37-280a-2529-39f751fca854', '{}', '69f00dda7cdc469f91c78020488ab380', '2020-08-30 22:40:18.214799', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-03 21:16:31.897677', 'ProcessingSettings', '流程设置', 0, NULL, '8c7e1739-c29c-edcb-b541-39f751de902a', 0, '/views/workflow/settings', '/views/workflow/settings', 'Operate.ProcessingSettings', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('de6dea49-731f-2a92-75ce-39f509369c1e', '{}', '20cd053ce02140e48636ed42dd0f7f20', '2020-05-09 09:25:48.584835', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-09 09:59:00.216529', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Icon', 'Icon图标', 0, '', 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/icons/index', '/views/icons/index', 'Infrastructure.Icon', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('e4abe794-13f9-c709-0c80-39f5092ce39f', '{}', 'ac66aa78de4246be942491b6efc78bcf', '2020-05-09 09:15:07.552637', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-12 21:38:21.469165', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Infrastructure', '基础设置', 2, 'international', NULL, 1, '/infrastructure', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('f7352ae7-6875-23e3-afe8-39f789a5d9f5', '{}', '18131777d7024f5abf7f3167201a1522', '2020-09-10 18:04:13.392838', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-10 20:05:23.286402', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Workflow', '工作流', 3, 'tree', NULL, 1, '/workflow', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');

-- ----------------------------
-- Table structure for au_module
-- ----------------------------
DROP TABLE IF EXISTS `au_module`;
CREATE TABLE `au_module`  (
  `ID` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '主键',
  `CreationTime` datetime NOT NULL COMMENT '创建时间',
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `LastModificationTime` datetime NULL DEFAULT NULL COMMENT '最后修改时间',
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '最后修改人',
  `IsDeleted` tinyint(1) NULL DEFAULT 0 COMMENT '是否已删除',
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '删除人',
  `DeletionTime` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `Name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模块名称',
  `Code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模块编码',
  `Rank` int(11) NOT NULL DEFAULT 1 COMMENT '排序',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of au_module
-- ----------------------------
INSERT INTO `au_module` VALUES ('0cbba6df-9078-4f4a-9b9c-8d6312d1b184', '2021-12-02 09:54:51', '66666666-6666-6666-6666-666666666666', '2021-12-02 11:28:32', '66666666-6666-6666-6666-666666666666', 0, '', NULL, '主数据', 'MasterDate', 100);
INSERT INTO `au_module` VALUES ('874bc286-f076-4591-929d-c2c2b34b6b86', '2021-12-02 10:41:20', '66666666-6666-6666-6666-666666666666', NULL, '', 0, '', NULL, '权限', 'Authorization', 99);
INSERT INTO `au_module` VALUES ('a4bfbd5a-4b33-4db6-934e-e1b2373f0744', '2021-12-02 10:55:25', '66666666-6666-6666-6666-666666666666', NULL, '', 1, '66666666-6666-6666-6666-666666666666', '2021-12-02 11:29:50', '测试', 'Test', 1000);

-- ----------------------------
-- Table structure for au_permission
-- ----------------------------
DROP TABLE IF EXISTS `au_permission`;
CREATE TABLE `au_permission`  (
  `Id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ExtraProperties` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `ConcurrencyStamp` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `CreationTime` datetime(6) NOT NULL,
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `LastModificationTime` datetime(6) NULL DEFAULT NULL,
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NOT NULL DEFAULT 0,
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeletionTime` datetime(6) NULL DEFAULT NULL,
  `MenuId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `PermissionGroup` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Policy` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Rank` int(11) NOT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of au_permission
-- ----------------------------
INSERT INTO `au_permission` VALUES ('15919b71-1d71-f23a-1759-39f543786b67', '{}', 'e6c082ad5dba4f249e10862c2b4ad30c', '2020-05-20 16:55:36.039229', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, 'afa5524e-b236-6607-1911-39f543786863', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('17f56317-da3a-57db-066b-39f55c3ebbf0', '{}', 'b610f669b9b4481d8f9ec56b1fc57f6e', '2020-05-25 12:23:05.968785', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-06-28 21:51:57.541222', '344f5ead-8a53-ea44-c25e-39f55c3eb97d', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('1861f3d0-de93-2a8a-4f3a-39f78a0fe93b', '{}', '65948b79f6ff4882b4c17e475ca6cf52', '2020-09-10 20:00:03.899087', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '3081e84b-656b-cbfd-5c8c-39f78a0fe83d', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('2ac09942-6016-1ff6-5b9d-39f53347d853', '{}', '4fe224ebf59745e29974c6b95aba17b2', '2020-05-17 13:28:37.203135', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'FunctionAuthorization', '功能授权', 'Operation.FunctionAuthorization', 2);
INSERT INTO `au_permission` VALUES ('2eaefbcf-650e-25e1-f31a-39f50b43e5c6', '{}', 'b052c3ca19e84bf295dc2af729bfe568', '2020-05-09 18:59:29.862441', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '5c01d798-9d1f-f088-9659-39f50b43e4a3', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('310a5777-ed92-ca37-a876-39f53345c243', '{}', '2bdd45209e1e4b59bd537ece5600d640', '2020-05-17 13:26:20.483446', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '216ad756-bdcc-abd3-c516-39f4ff08b978', 'Operation', 'Create', '新增', 'Operation.Create', 1);
INSERT INTO `au_permission` VALUES ('346f6545-bad6-4273-8fcf-ad304402e8c3', NULL, NULL, '2022-02-12 21:04:05.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, 'ce6c3946-3da7-493a-a5c2-1d5ad4444aee', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('3a9f83cc-0ba6-fda7-9848-39f56645da88', '{}', 'e0b60d0719f94426980312274595c5aa', '2020-05-27 11:07:04.712364', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-27 11:28:31.715106', 'b5047074-9a0c-8721-b2a0-39f56645d7b0', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('3b386f9a-2efe-1f67-2e78-39f533482e9e', '{}', '1b741e0d826d425289134f76feee4faa', '2020-05-17 13:28:59.294866', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'Edit', '编辑', 'Operation.Edit', 3);
INSERT INTO `au_permission` VALUES ('430ec3ec-ba7a-8d72-9346-39f53348fa1b', '{}', '34fcfef25549427b9e8b61022049cb17', '2020-05-17 13:29:51.387740', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'Delete', '删除', 'Operation.Delete', 5);
INSERT INTO `au_permission` VALUES ('4521f526-9e2c-f208-5e53-39f533461812', '{}', '4fcbf073615848faa32b4fc7cc5a14a7', '2020-05-17 13:26:42.450320', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '216ad756-bdcc-abd3-c516-39f4ff08b978', 'Operation', 'Edit', '编辑', 'Operation.Edit', 2);
INSERT INTO `au_permission` VALUES ('5d4ed6df-6eb5-fad4-f953-39f53346635f', '{}', '3e1d85a3729842ec960bc6ad08102941', '2020-05-17 13:27:01.727705', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '216ad756-bdcc-abd3-c516-39f4ff08b978', 'Operation', 'Delete', '删除', 'Operation.Delete', 3);
INSERT INTO `au_permission` VALUES ('5ee2d035-47ab-482b-805b-1bf4ffa4d4f6', '{}', 'b052c3ca19e84bf295dc2af729bfe568', '2020-05-09 18:59:29.862441', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '216ad756-bdcc-abd3-c516-39f4ff08b978', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('5f092eaf-a7ca-5cad-8e3a-39f5109e472d', '{}', '9fe9ace282de45259e91a1db397cbdac', '2020-05-10 19:56:19.117957', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('77520662-a275-fa39-195f-39f55c39124e', '{}', 'cd90668e5dce48f8b78b7b37e2f0f861', '2020-05-25 12:16:54.862505', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '93352161-b3fb-f99e-5cd9-39f55c390fec', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('79646120-b768-6b63-3724-39f5334426fc', '{}', '3cb693c3f3f44b65940668233bfb1057', '2020-05-17 13:24:35.197675', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '332f6c3f-40cc-2908-99ef-39f4e6b0dfec', 'Operation', 'Create', '新增', 'Operation.Create', 3);
INSERT INTO `au_permission` VALUES ('8eded392-4793-229f-b9a1-39f50936aba8', '{}', 'a88ea79551f641ba960b4c4731aa3c71', '2020-05-09 09:25:48.585026', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, 'de6dea49-731f-2a92-75ce-39f509369c1e', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('a6182e8c-ee39-8954-fd0d-39f53348b89c', '{}', '11286c1367794c85bc6e3b6a4e58037d', '2020-05-17 13:29:34.620458', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-14 16:54:09.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'EditStatus', '启用/禁用', 'Operation.EditStatus', 4);
INSERT INTO `au_permission` VALUES ('ad82e160-208e-f7ca-6b72-39f519df2863', '{}', '59bbf51674e44f69b32c8d4c6ce0a1af', '2020-05-12 15:03:46.019618', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-12 15:04:02.407300', 'de6dea49-731f-2a92-75ce-39f509369c1e', 'View', 'View', '查看', 'View.View', 1);
INSERT INTO `au_permission` VALUES ('ae60fc43-dbc0-3501-9de1-39f53347a8af', '{}', '5efbf3c2438e4d07a0f5374dd3594e98', '2020-05-17 13:28:25.007350', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'Create', '新增', 'Operation.Create', 1);
INSERT INTO `au_permission` VALUES ('af543e5e-22de-38a8-6b0d-39f519c854ab', '{}', '5674cba074174f5dbd5f9e2813d26944', '2020-05-12 14:38:50.028212', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-13 08:45:01.338292', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '332f6c3f-40cc-2908-99ef-39f4e6b0dfec', 'Operation', 'Edit', '编辑', 'Operation.Edit', 1);
INSERT INTO `au_permission` VALUES ('b58a86c9-041b-4db8-af39-5723fc955375', '{}', 'b052c3ca19e84bf295dc2af729bfe568', '2020-05-09 18:59:29.862441', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '332f6c3f-40cc-2908-99ef-39f4e6b0dfec', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('b5cb10e2-f1b8-36b4-679a-39f519e1b9d1', '{}', '1f7c1e4fd0ea4fa6828f98aad498ee0e', '2020-05-12 15:06:34.321069', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '332f6c3f-40cc-2908-99ef-39f4e6b0dfec', 'Operation', 'Delete', '删除', 'Operation.Delete', 2);
INSERT INTO `au_permission` VALUES ('d98eadfd-07ce-142b-2778-39f5343364fd', '{}', 'b869eb86df7e428691dd86b8289a3c56', '2020-05-17 17:45:54.173146', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '753ce348-567c-4e19-22f3-39f534336219', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('dbe722c5-5e04-a1de-ab54-39f4ff08b979', '{}', 'dbd9ad330d224390a5f334c9d7c12751', '2020-05-07 09:59:25.305726', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '00000000-0000-0000-0000-000000000000', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('eab51e5d-4d6e-4361-8ab1-39f751fca927', '{}', '746429c92fa4465c8b2f600a4e85834b', '2020-08-30 22:40:18.215832', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-03 21:16:31.844523', 'd4ff0631-df37-280a-2529-39f751fca854', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('f57df6ef-4e2e-b962-ee82-39f50a959c6a', '{}', 'b10d858b19534a1ba2da9bc7def9dafb', '2020-05-09 15:49:07.818781', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-10 18:48:00.953912', 'a6723edb-74e4-3053-c16b-39f50a959be0', 'View', 'Access', '页面访问', 'View.Access', 0);

-- ----------------------------
-- Table structure for au_role
-- ----------------------------
DROP TABLE IF EXISTS `au_role`;
CREATE TABLE `au_role`  (
  `Id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ExtraProperties` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `ConcurrencyStamp` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `CreationTime` datetime(6) NOT NULL,
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `LastModificationTime` datetime(6) NULL DEFAULT NULL,
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NOT NULL DEFAULT 0,
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeletionTime` datetime(6) NULL DEFAULT NULL,
  `TenantId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Code` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `RoleStatus` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of au_role
-- ----------------------------
INSERT INTO `au_role` VALUES ('09019ba5-e2b9-0a70-7181-39f751a51b5d', '{}', 'a908341fb4e24c3aa5dfd216b7c63431', '2020-08-30 21:04:40.285231', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-12 20:04:21.000000', '88888888-8888-8888-8888-888888888888', NULL, 'twe', NULL, 'Enable');
INSERT INTO `au_role` VALUES ('116f4336-0793-f5a3-aa44-39f751a5491a', '{}', '235c3994dc904e7ea826538d46c47592', '2020-08-30 21:04:51.994427', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-12 20:07:39.000000', '88888888-8888-8888-8888-888888888888', NULL, 'uert', NULL, 'Enable');
INSERT INTO `au_role` VALUES ('1212ea92-63b5-54d6-b20f-39f751a55ef7', '{}', '2afaf9eb1bac4b9d808b8a2aad4a7c86', '2020-08-30 21:04:57.591684', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-13 22:06:40.000000', '88888888-8888-8888-8888-888888888888', NULL, 'yyyyy', NULL, 'Enable');
INSERT INTO `au_role` VALUES ('243c22de-85a5-b022-99f0-39f751a538dd', '{}', 'ac1b2889143e470fa84159f14e102c04', '2020-08-30 21:04:47.837126', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-13 22:06:42.000000', '88888888-8888-8888-8888-888888888888', NULL, 'eyu', NULL, 'Enable');
INSERT INTO `au_role` VALUES ('2e6de576-0e5a-40b3-a273-6692ab313fab', NULL, NULL, '2022-02-14 16:57:05.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '', '测试角色', '测试专用', 'Enable');
INSERT INTO `au_role` VALUES ('3db40220-0c5a-3bdf-8d0e-39f751a4fb64', '{}', '3a3ccda68d134f34ac5bd3dea6207c4e', '2020-08-30 21:04:32.100980', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-13 22:06:45.000000', '88888888-8888-8888-8888-888888888888', NULL, '1254', NULL, 'Enable');
INSERT INTO `au_role` VALUES ('42d2f479-0b3f-4c29-4419-39f51afdf387', '{}', 'ba12f9e866f24a31b1242aff33cd0e1a', '2020-05-12 20:17:01.327669', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', NULL, '系统管理员', '所有菜单权限', 'Enable');
INSERT INTO `au_role` VALUES ('474df367-c814-21db-3447-39f751a52acd', '{}', '50e2030eab734de783da7b085004f70a', '2020-08-30 21:04:44.237923', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-13 22:06:50.000000', '88888888-8888-8888-8888-888888888888', NULL, 'yey', NULL, 'Enable');
INSERT INTO `au_role` VALUES ('4c86b621-61a2-5d15-aaf5-39f51aff97b2', '{}', '420cafa57a45474d9f4d3cf63afd7b4d', '2020-05-12 20:18:48.882198', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-06-03 15:28:07.554096', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', NULL, '普通用户', '普通的功能权限11', 'Enable');
INSERT INTO `au_role` VALUES ('5d7e5def-c0d6-3625-f611-39f751a4c478', '{}', '132d9ee54d894cfaacbb75c48aa79d91', '2020-08-30 21:04:18.041361', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-13 22:06:57.000000', '88888888-8888-8888-8888-888888888888', NULL, 'sss', '1', 'Enable');
INSERT INTO `au_role` VALUES ('7212c178-e2e5-9423-59c4-39f751a4d9eb', '{}', 'bafa9250409d4bf0b8fe2a2f8eda5aac', '2020-08-30 21:04:23.531762', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-13 22:06:59.000000', '88888888-8888-8888-8888-888888888888', NULL, '122', '1', 'Enable');
INSERT INTO `au_role` VALUES ('ccd08f86-c592-35f3-d33d-39f751a50c6e', '{}', 'eb509b02f48d4409b80b2c65d7665869', '2020-08-30 21:04:36.462046', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-13 22:07:02.000000', '88888888-8888-8888-8888-888888888888', NULL, 'wer', NULL, 'Enable');
INSERT INTO `au_role` VALUES ('eabb4b95-b94e-d0e3-932a-39f751a4ea07', '{}', '2033cd8e95874947a5f7e4a4efe2d308', '2020-08-30 21:04:27.655694', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-13 22:07:05.000000', '88888888-8888-8888-8888-888888888888', NULL, '124', NULL, 'Enable');

-- ----------------------------
-- Table structure for au_rolepermission
-- ----------------------------
DROP TABLE IF EXISTS `au_rolepermission`;
CREATE TABLE `au_rolepermission`  (
  `Id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `CreationTime` datetime(6) NOT NULL,
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `LastModificationTime` datetime(6) NULL DEFAULT NULL,
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NOT NULL DEFAULT 0,
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeletionTime` datetime(6) NULL DEFAULT NULL,
  `TenantId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `RoleId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `PermissionId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `IsGranted` tinyint(100) NOT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of au_rolepermission
-- ----------------------------
INSERT INTO `au_rolepermission` VALUES ('01709331-81ef-726d-9df3-39f5328c7847', '2020-05-17 10:03:57.417414', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-17 10:09:24.165480', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '42d2f479-0b3f-4c29-4419-39f51afdf387', 'af543e5e-22de-38a8-6b0d-39f519c854ab', 1);
INSERT INTO `au_rolepermission` VALUES ('0dec0e46-df0c-4b3f-af75-3b67514c1aca', '2022-02-15 23:30:55.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', '2eaefbcf-650e-25e1-f31a-39f50b43e5c6', 1);
INSERT INTO `au_rolepermission` VALUES ('1a5ba178-d480-ea54-2fdc-39f5328c7849', '2020-05-17 10:03:57.417479', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-17 10:09:24.165494', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '42d2f479-0b3f-4c29-4419-39f51afdf387', 'b58a86c9-041b-4db8-af39-5723fc955375', 1);
INSERT INTO `au_rolepermission` VALUES ('1d525bcc-c3fc-6bd7-9cf3-39f53291d139', '2020-05-17 10:09:47.833821', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '4c86b621-61a2-5d15-aaf5-39f51aff97b2', '8eded392-4793-229f-b9a1-39f50936aba8', 1);
INSERT INTO `au_rolepermission` VALUES ('1df7a45d-9b64-771d-e046-39f5328c784b', '2020-05-17 10:03:57.417482', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-17 10:09:24.165511', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '42d2f479-0b3f-4c29-4419-39f51afdf387', 'b5cb10e2-f1b8-36b4-679a-39f519e1b9d1', 1);
INSERT INTO `au_rolepermission` VALUES ('1e70aac4-096c-a547-ef92-39f5328c784f', '2020-05-17 10:03:57.417488', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-17 10:09:24.165547', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '42d2f479-0b3f-4c29-4419-39f51afdf387', '8eded392-4793-229f-b9a1-39f50936aba8', 1);
INSERT INTO `au_rolepermission` VALUES ('59be8e83-0d92-a492-949f-39f677db3dcd', '2020-07-19 14:06:34.980419', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '4c86b621-61a2-5d15-aaf5-39f51aff97b2', '15919b71-1d71-f23a-1759-39f543786b67', 1);
INSERT INTO `au_rolepermission` VALUES ('5faba319-cefd-463c-9c21-3361e6ca8908', '2022-02-15 23:31:34.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', '5f092eaf-a7ca-5cad-8e3a-39f5109e472d', 1);
INSERT INTO `au_rolepermission` VALUES ('6bca577b-7327-4bde-adb2-c23a7b4dcc3a', '2022-02-15 23:32:50.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', '1861f3d0-de93-2a8a-4f3a-39f78a0fe93b', 1);
INSERT INTO `au_rolepermission` VALUES ('7601c7f8-93cc-d6f8-51c6-39f5328c784d', '2020-05-17 10:03:57.417486', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-17 10:09:24.165538', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '42d2f479-0b3f-4c29-4419-39f51afdf387', '5f092eaf-a7ca-5cad-8e3a-39f5109e472d', 1);
INSERT INTO `au_rolepermission` VALUES ('84f88bf2-28cc-4a5b-bc4f-85be4281c3c5', '2022-02-15 23:31:34.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', '15919b71-1d71-f23a-1759-39f543786b67', 1);
INSERT INTO `au_rolepermission` VALUES ('90b9b0e9-1477-6202-6d37-39f5328c7820', '2020-05-17 10:03:57.416350', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-17 10:09:24.165444', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '42d2f479-0b3f-4c29-4419-39f51afdf387', '2eaefbcf-650e-25e1-f31a-39f50b43e5c6', 1);
INSERT INTO `au_rolepermission` VALUES ('94a06490-f1a1-4004-9043-0804b87f98ff', '2022-02-15 23:31:34.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-16 10:52:42.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', '77520662-a275-fa39-195f-39f55c39124e', 1);
INSERT INTO `au_rolepermission` VALUES ('98e4e705-4d94-4f24-a2a8-6a332c1efd0e', '2022-02-15 23:31:34.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', 'd98eadfd-07ce-142b-2778-39f5343364fd', 1);
INSERT INTO `au_rolepermission` VALUES ('c727e6ef-e454-40bc-ae49-48732b966810', '2022-02-15 23:31:34.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', 'b58a86c9-041b-4db8-af39-5723fc955375', 1);
INSERT INTO `au_rolepermission` VALUES ('c98a8515-158a-e962-383d-39f5328c784c', '2020-05-17 10:03:57.417484', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-17 10:09:24.165529', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '42d2f479-0b3f-4c29-4419-39f51afdf387', '5ee2d035-47ab-482b-805b-1bf4ffa4d4f6', 1);
INSERT INTO `au_rolepermission` VALUES ('d5038274-ab87-84c3-378e-39f53291d138', '2020-05-17 10:09:47.833798', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', '4c86b621-61a2-5d15-aaf5-39f51aff97b2', '2eaefbcf-650e-25e1-f31a-39f50b43e5c6', 1);
INSERT INTO `au_rolepermission` VALUES ('d99874f6-c23e-4833-9def-ee7c5f576a90', '2022-02-15 23:31:34.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', '8eded392-4793-229f-b9a1-39f50936aba8', 1);
INSERT INTO `au_rolepermission` VALUES ('ff43c403-9ad0-45c9-a8ea-a515471b46d3', '2022-02-15 23:31:34.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '2e6de576-0e5a-40b3-a273-6692ab313fab', '5ee2d035-47ab-482b-805b-1bf4ffa4d4f6', 1);

-- ----------------------------
-- Table structure for au_userdataprivilege
-- ----------------------------
DROP TABLE IF EXISTS `au_userdataprivilege`;
CREATE TABLE `au_userdataprivilege`  (
  `Id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `CreationTime` datetime(6) NOT NULL,
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `LastModificationTime` datetime(6) NULL DEFAULT NULL,
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NOT NULL DEFAULT 0,
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeletionTime` datetime(6) NULL DEFAULT NULL,
  `TenantId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `UserId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `RoleId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of au_userdataprivilege
-- ----------------------------
INSERT INTO `au_userdataprivilege` VALUES ('099a3b2b-e371-437e-b944-80e6e2e6c89f', '2022-02-16 23:29:52.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-18 00:12:27.000000', '88888888-8888-8888-8888-888888888888', '20ec0e13-5ca5-4083-b788-6cd8492a7170', '4c86b621-61a2-5d15-aaf5-39f51aff97b2');
INSERT INTO `au_userdataprivilege` VALUES ('1cf0d250-29d7-4dcd-a98b-e7c557bd9489', '2022-02-18 00:31:08.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '8cc9e821-99bb-4b60-9f06-77b5eb11cc76', '2e6de576-0e5a-40b3-a273-6692ab313fab');
INSERT INTO `au_userdataprivilege` VALUES ('2984ed70-0192-c99b-c60f-39f677502afe', '2020-07-19 11:34:40.734247', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-18 00:12:27.000000', '88888888-8888-8888-8888-888888888888', '6a95b42c-9fce-08ac-6151-39f677502834', '4c86b621-61a2-5d15-aaf5-39f51aff97b2');
INSERT INTO `au_userdataprivilege` VALUES ('5b371d86-fc95-04ef-2e82-39f534eed301', '2020-05-17 21:10:38.222918', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-18 00:12:27.000000', '88888888-8888-8888-8888-888888888888', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '42d2f479-0b3f-4c29-4419-39f51afdf387');
INSERT INTO `au_userdataprivilege` VALUES ('5fb32b0b-2d3c-69d3-45bc-39f66019912b', '2020-07-14 23:23:46.434826', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-18 00:12:27.000000', '88888888-8888-8888-8888-888888888888', '08d75a27-7ac1-6bff-0e37-4d2115c732d5', '4c86b621-61a2-5d15-aaf5-39f51aff97b2');
INSERT INTO `au_userdataprivilege` VALUES ('91c003c0-db33-3143-b19d-39f6650335e4', '2020-07-15 22:17:27.365110', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-07-15 22:46:58.409329', '88888888-8888-8888-8888-888888888888', '08d75a27-7ac1-6bff-0e37-4d2115c732d5', '4c86b621-61a2-5d15-aaf5-39f51aff97b2');
INSERT INTO `au_userdataprivilege` VALUES ('c1991749-9f5e-4ac1-af84-05162b67465c', '2022-02-18 00:31:05.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '20ec0e13-5ca5-4083-b788-6cd8492a7170', '4c86b621-61a2-5d15-aaf5-39f51aff97b2');
INSERT INTO `au_userdataprivilege` VALUES ('c8da2eca-653b-4b9f-9d60-ec111adc4c9f', '2022-02-18 00:12:40.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-18 00:13:01.000000', '88888888-8888-8888-8888-888888888888', '20ec0e13-5ca5-4083-b788-6cd8492a7170', '2e6de576-0e5a-40b3-a273-6692ab313fab');
INSERT INTO `au_userdataprivilege` VALUES ('ceba2c3e-3455-4e01-9f86-8ade55ec1d39', '2022-02-18 00:30:15.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '88888888-8888-8888-8888-888888888888', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '42d2f479-0b3f-4c29-4419-39f51afdf387');
INSERT INTO `au_userdataprivilege` VALUES ('f97efd0f-7dbf-43b9-a900-4c795af55e60', '2022-02-18 00:21:09.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-18 00:28:15.000000', '88888888-8888-8888-8888-888888888888', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '42d2f479-0b3f-4c29-4419-39f51afdf387');

-- ----------------------------
-- Table structure for sys_option
-- ----------------------------
DROP TABLE IF EXISTS `sys_option`;
CREATE TABLE `sys_option`  (
  `Id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ExtraProperties` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `ConcurrencyStamp` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `CreationTime` datetime(6) NOT NULL,
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `LastModificationTime` datetime(6) NULL DEFAULT NULL,
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NOT NULL DEFAULT 0,
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeletionTime` datetime(6) NULL DEFAULT NULL,
  `GroupName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `GroupKey` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `EnumCode` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `EnumName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `EnumLabel` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Levels` int(11) NOT NULL,
  `Orders` int(11) NOT NULL,
  `IsHide` tinyint(1) NOT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_option
-- ----------------------------

-- ----------------------------
-- Table structure for sys_uploadfile
-- ----------------------------
DROP TABLE IF EXISTS `sys_uploadfile`;
CREATE TABLE `sys_uploadfile`  (
  `Id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ExtraProperties` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `ConcurrencyStamp` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `CreationTime` datetime(6) NOT NULL,
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `LastModificationTime` datetime(6) NULL DEFAULT NULL,
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NOT NULL DEFAULT 0,
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeletionTime` datetime(6) NULL DEFAULT NULL,
  `TenantId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `FileName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `FilePath` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `FileSize` bigint(20) NOT NULL,
  `FileType` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Status` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `IX_Sys_UploadFile_FilePath`(`FilePath`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_uploadfile
-- ----------------------------
INSERT INTO `sys_uploadfile` VALUES ('5055599d-6d4c-a54c-8e11-3a00ee222569', '{\"QiNiu\":{\"Code\":200,\"Text\":\"{\\\"hash\\\":\\\"FqlccOuZrou2QGl_KKqC8L-bMEGN\\\",\\\"key\\\":\\\"5632d084-f639-3003-6c54-3a00ee222042.png\\\"}\",\"Data\":null,\"RefCode\":200,\"RefText\":\"[2021-12-21 15:47:05.1133] [FormUpload] Uploaded: \\\"#STREAM#\\\" ==> \\\"5632d084-f639-3003-6c54-3a00ee222042.png\\\"\\n\",\"RefInfo\":{\"Content-Type\":\"application/json\",\"Content-Length\":\"88\",\"Server\":\"openresty\",\"Date\":\"Tue, 21 Dec 2021 07:47:04 GMT\",\"Connection\":\"keep-alive\",\"Access-Control-Allow-Headers\":\"X-File-Name, X-File-Type, X-File-Size\",\"Access-Control-Allow-Methods\":\"OPTIONS, HEAD, POST\",\"Access-Control-Allow-Origin\":\"*\",\"Access-Control-Expose-Headers\":\"X-Log, X-Reqid\",\"Access-Control-Max-Age\":\"2592000\",\"Cache-Control\":\"no-store, must-revalidate, no-cache\",\"Pragma\":\"no-cache\",\"X-Content-Type-Options\":\"nosniff\",\"X-Reqid\":\"R38AAAAWG2aLtcIW\",\"X-Svr\":\"UP\",\"X-Log\":\"X-Log\"}}}', '8b2f394bfb5a473495b4ae7d1bd33d04', '2021-12-21 15:47:06.218086', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', 'map.png', '5632d084-f639-3003-6c54-3a00ee222042.png', 126703, 'Image', 'NotEnabled');
INSERT INTO `sys_uploadfile` VALUES ('9eeec37f-5a77-6ecf-e3dc-3a00ee8a1a67', '{\"QiNiu\":{\"Code\":200,\"Text\":\"{\\\"hash\\\":\\\"FsTJrejT7K0EYbPRamdCMBugr0pP\\\",\\\"key\\\":\\\"9b15509e-fdfd-926b-f58c-3a00ee8a18b2.png\\\"}\",\"Data\":null,\"RefCode\":200,\"RefText\":\"[2021-12-21 17:40:38.9214] [FormUpload] Uploaded: \\\"#STREAM#\\\" ==> \\\"9b15509e-fdfd-926b-f58c-3a00ee8a18b2.png\\\"\\n\",\"RefInfo\":{\"Content-Type\":\"application/json\",\"Content-Length\":\"88\",\"Server\":\"openresty\",\"Date\":\"Tue, 21 Dec 2021 09:40:38 GMT\",\"Connection\":\"keep-alive\",\"Access-Control-Allow-Headers\":\"X-File-Name, X-File-Type, X-File-Size\",\"Access-Control-Allow-Methods\":\"OPTIONS, HEAD, POST\",\"Access-Control-Allow-Origin\":\"*\",\"Access-Control-Expose-Headers\":\"X-Log, X-Reqid\",\"Access-Control-Max-Age\":\"2592000\",\"Cache-Control\":\"no-store, must-revalidate, no-cache\",\"Pragma\":\"no-cache\",\"X-Content-Type-Options\":\"nosniff\",\"X-Reqid\":\"ycgAAAAgTd29u8IW\",\"X-Svr\":\"UP\",\"X-Log\":\"X-Log\"}}}', 'f93da1628e2c4d83a9c11c55e13bf241', '2021-12-21 17:40:39.143650', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', 'ser1.png', '9b15509e-fdfd-926b-f58c-3a00ee8a18b2.png', 23986, 'Image', 'NotEnabled');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `Id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ExtraProperties` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `ConcurrencyStamp` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `CreationTime` datetime(6) NOT NULL,
  `CreatorId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `LastModificationTime` datetime(6) NULL DEFAULT NULL,
  `LastModifierId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NOT NULL DEFAULT 0,
  `DeleterId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `DeletionTime` datetime(6) NULL DEFAULT NULL,
  `TenantId` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Status` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `UserName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `Password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `LastIp` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `LastLogDate` datetime(6) NULL DEFAULT NULL,
  `HeadImgUrl` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Sex` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Mobile` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `IsTenantAdmin` tinyint(1) NOT NULL DEFAULT 0,
  `LoginName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('08d74d8e-4e45-c152-f0b2-22013a71bdde', '{}', 'a34410fbe5d64fee865e7cb767b77bd4', '2019-10-10 22:29:49.163310', '08d72959-04b8-f620-f176-e9cdf31aab96', '2022-02-18 00:30:15.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '88888888-8888-8888-8888-888888888888', 'Enable', 'admin', '$2a$10$Yul/8nQGrVVgzAAaFwgbb.BWn4meAzW5vlOyvO1dY.BoTGMXWHMVC', NULL, NULL, '/group1/hfmall/20200719/12/00/7/89076334ea9d02a295734756fc1d7843.png', 'Secrecy', '13560720489', '455650119@qq.com', 1, 'admin');
INSERT INTO `sys_user` VALUES ('08d75a27-7ac1-6bff-0e37-4d2115c732d5', '{}', '910510683e5a4b56a0d86e5eedfa8f27', '2019-10-26 23:16:30.666959', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-07-15 22:47:25.237969', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-12 19:04:38.000000', '88888888-8888-8888-8888-888888888888', 'Enable', 'bryan', '$2a$10$Yul/8nQGrVVgzAAaFwgbb.BWn4meAzW5vlOyvO1dY.BoTGMXWHMVC', NULL, NULL, '/group1/hfmall/20200714/23/23/7/9ed2396d606e2f0a688796ba00680f41.png', 'Male', '123456789', 'string', 0, 'bryan');
INSERT INTO `sys_user` VALUES ('20ec0e13-5ca5-4083-b788-6cd8492a7170', NULL, NULL, '2022-02-12 19:08:51.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-18 00:31:05.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, '88888888-8888-8888-8888-888888888888', 'Enable', '测试账号', '$2a$10$3hKeXt2bY3eBq4Pl/7Sgt.5WScdl.NA5dRaFrbz6I.5GWsmXqm12G', NULL, NULL, '', NULL, NULL, '', 0, 'demo1');
INSERT INTO `sys_user` VALUES ('6a95b42c-9fce-08ac-6151-39f677502834', '{}', '25d202302a7d4c51aaee611e3dc5c119', '2020-07-19 11:34:40.032578', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, NULL, NULL, '88888888-8888-8888-8888-888888888888', 'Enable', '测试账号', '$2a$10$Yul/8nQGrVVgzAAaFwgbb.BWn4meAzW5vlOyvO1dY.BoTGMXWHMVC', NULL, NULL, '/group1/hfmall/20200719/11/33/7/cceaa0bbe6dae640a4570ce297dec9ad.png', 'Secrecy', NULL, NULL, 0, 'test');
INSERT INTO `sys_user` VALUES ('8cc9e821-99bb-4b60-9f06-77b5eb11cc76', NULL, NULL, '2022-01-29 17:12:53.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-18 00:31:08.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, '88888888-8888-8888-8888-888888888888', 'Enable', '测试账号', '$2a$10$a93JbJSvMTp9g6hDrXESCe8wjiolOg5pqMqHF5alFCND/GokoVAK6', NULL, NULL, 'http://dummyimage.com/100x100', NULL, NULL, '', 0, 'demo');

SET FOREIGN_KEY_CHECKS = 1;
