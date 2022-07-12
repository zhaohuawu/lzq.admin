/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.10.103-13306
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : 192.168.10.103:13306
 Source Schema         : lzq_admin_test

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 12/07/2022 21:12:13
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
  `Icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
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
INSERT INTO `au_menu` VALUES ('041d4d04-18db-469f-b44e-8f916747b143', NULL, NULL, '2022-05-21 19:13:25.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-05-21 21:33:31.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'LoginLogList', '登录日志', 2, '', '942a2853-8f9f-4214-aa7f-53c363947d03', 0, '/auditlog-action/loginlog-list', '/views/auditlog-action/loginlog-list', 'AuditLog.LoginLogList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('075ccc2d-35af-0455-2e9c-39f50af7b4ad', '{}', '6092df602127449f8d298f7cf4812404', '2020-05-09 17:36:16.728482', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-09 18:23:37.651060', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-12 11:26:11.410298', 'TenantManagement', '租户管理', 2, 'peoples', NULL, 1, '/tenantmanagement', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('0d37c57f-995a-45ce-a8fb-5555c519f5cb', NULL, NULL, '2022-05-23 21:41:49.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-05-23 21:41:55.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'Organization', '部门管理', 3, '', 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/organization/index', '/views/organization/index', 'Infrastructure.Organization', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('1216c0e1-efd7-4071-b5bb-f3689365d911', NULL, NULL, '2022-06-08 11:56:45.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, 'Workflow', '流程管理', 3, 'cascader', '', 1, '/workflow', '', '', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', '{}', '1613b6da5f774156b08d709500614f85', '2020-05-10 19:56:19.113957', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-06-08 21:30:47.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'RoleList', '角色管理', 2, '', 'cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', 0, '/authorization/role-list', '/views/authorization/role-list', 'Authorization.RoleList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('1aac4140-ef3d-4467-9de0-5a63ddfe9287', NULL, NULL, '2022-07-10 09:28:44.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-07-10 09:30:08.000000', 'SystemConfig', '系统配置', 99, '', 'e4abe794-13f9-c709-0c80-39f5092ce39f', 1, '/', '', '', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('216ad756-bdcc-abd3-c516-39f4ff08b978', '{}', 'ffc3ecdf9e6c44e9b8c55289fa78a460', '2020-05-07 09:59:25.305045', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-06-08 21:39:55.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'PermissionList', '操作权限', 1, '', 'cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', 0, '/authorization/permission-list', '/views/authorization/permission-list', 'Authorization.PermissionList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('3081e84b-656b-cbfd-5c8c-39f78a0fe83d', '{}', '9b22ec2aac124158b576a858788de36e', '2020-09-10 20:00:03.898518', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-21 20:44:41.485719', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'ProcessSettings', '流程设置', 1, NULL, 'f7352ae7-6875-23e3-afe8-39f789a5d9f5', 0, '/workflow/approvaltype-list', '/workflow/approvaltype-list', 'Workflow.ProcessSettings', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('332f6c3f-40cc-2908-99ef-39f4e6b0dfec', '{}', '450ef3e0fc2a431792c4f8b8ddae3f8c', '2020-05-02 16:32:34.797764', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-06-08 21:39:46.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'MenuList', '菜单管理', 0, '', 'cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', 0, '/authorization/menu-list', '/views/authorization/menu-list', 'Authorization.MenuList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('33c763ff-62b2-a190-a3ec-39f50b4e28cf', '{}', '36cb15fe1bd442388517ec2e3700b8c9', '2020-05-09 19:10:42.460452', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-03-09 22:47:33.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Dashboard', '首页容器', 0, 'dashboard', '', 1, '/', '', '', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('344f5ead-8a53-ea44-c25e-39f55c3eb97d', '{}', '334a02f02033400ca2ce3050f38c4fe8', '2020-05-25 12:23:05.967979', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-06-03 16:18:55.038789', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-06-28 21:51:57.597826', 'Depts', '部门管理1', 2, NULL, 'cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', 0, '/dept/index', '/views/dept/index', 'Authorization.Depts', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('56410b19-72bb-42c2-a53b-b4082a58bdff', NULL, NULL, '2022-06-08 11:59:02.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-06-08 14:03:43.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'Designer', '流程设计器', 1, '', '1216c0e1-efd7-4071-b5bb-f3689365d911', 0, '/workflow/designer', '/views/workflow/designer', 'Workflow.Designer', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('59fc719d-bbfd-4059-a20c-6672dba08b06', NULL, NULL, '2022-05-05 00:07:14.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-07-10 17:01:36.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'SystemConfig', '系统配置', 98, '', 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/systemconfig/index', '/views/systemconfig/index', 'Infrastructure.SystemConfig', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('5c01d798-9d1f-f088-9659-39f50b43e4a3', '{}', 'b09bf397a8ae42b0a3ca3046e3af94f5', '2020-05-09 18:59:29.860932', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-23 16:38:59.057788', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Home', '首页', 0, '', '33c763ff-62b2-a190-a3ec-39f50b4e28cf', 0, 'dashboard', '/views/dashboard/index', 'Dashboard.HomePage', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('5dd65781-eeca-4e8f-b545-4abdf0e1d8c5', NULL, NULL, '2022-06-08 14:00:58.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-06-08 14:03:49.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'ProcessSettings', '流程设计', 1, '', '1216c0e1-efd7-4071-b5bb-f3689365d911', 0, '/workflow/index', '/views/workflow/index', 'Workflow.ProcessSettings', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('753ce348-567c-4e19-22f3-39f534336219', '{}', '70047beb29ab4358bec1c4b354900163', '2020-05-17 17:45:54.171447', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-06-07 23:40:22.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'SysUserList', '用户管理', 1, '', 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/sysuser/sysuser-list', '/views/sysuser/sysuser-list', 'Infrastructure.SysUserList', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('84cfd384-8062-434d-912a-4c39c97a844b', NULL, NULL, '2022-07-10 11:20:50.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-07-10 19:19:18.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'SystemDictionary', '字典管理', 1, '', 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/systemdict/index', '/views/systemdict/index', 'Infrastructure.SystemDictionary', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('8c7e1739-c29c-edcb-b541-39f751de902a', '{}', 'a473cc5deed24f619e5af6b06811e980', '2020-08-30 22:07:25.853392', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-03 22:48:24.944178', 'Workflow', '工作流', 3, 'tree', NULL, 1, 'workflow', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('93352161-b3fb-f99e-5cd9-39f55c390fec', '{}', 'eb0df579ceb84e66be9e40e05ac9afaf', '2020-05-25 12:16:54.861810', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-03-19 20:03:04.000000', 'Settings', '配置设置', 3, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/settings', '/views/settings/index', 'Infrastructure.Settings', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('942a2853-8f9f-4214-aa7f-53c363947d03', NULL, NULL, '2022-05-21 19:02:52.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-05-21 21:34:22.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'AuditLog', '日志管理', 4, 'log', '', 1, '/auditlog', '', '', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('a6723edb-74e4-3053-c16b-39f50a959be0', '{}', 'f63b12e3a9304fc19c3863c13576685f', '2020-05-09 15:49:07.818682', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-10 18:15:04.200070', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-10 18:48:01.005610', 'Logs', '日志k', 1, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, 'logs', '/view/logs/index', 'Infrastructure.Logs', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('afa5524e-b236-6607-1911-39f543786863', '{}', 'ce0e14169f124ec29b0bd14517b6fcdb', '2020-05-20 16:55:36.037295', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-27 15:15:50.422349', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-03-19 20:03:01.000000', 'Log', '日志', 2, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/log/index', '/view/log/index', 'Infrastructure.Log', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('b5047074-9a0c-8721-b2a0-39f56645d7b0', '{}', '3c955a4511574eb4b6989c4c4ebeed7a', '2020-05-27 11:07:04.710383', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-27 11:28:31.773987', 'Areas', '省市区管理', 4, NULL, 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/areas/index', '/views/areas/index', 'Infrastructure.Areas', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('cd788d26-3fbd-bb4a-10d8-39f4f5159fa8', '{}', 'bdeda6e5e9784732999556c5425b5f56', '2020-05-05 11:37:18.505680', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-03-09 22:46:02.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Authorization', '权限管理', 1, 'lock', '', 1, '/authorization', '', '', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('ce6c3946-3da7-493a-a5c2-1d5ad4444aee', NULL, NULL, '2022-02-12 21:04:05.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, 'ceshi', 'ceshi', 1, '', 'f7352ae7-6875-23e3-afe8-39f789a5d9f5', 0, '/', '/', 'Workflow.ceshi', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('d4ff0631-df37-280a-2529-39f751fca854', '{}', '69f00dda7cdc469f91c78020488ab380', '2020-08-30 22:40:18.214799', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-03 21:16:31.897677', 'ProcessingSettings', '流程设置', 0, NULL, '8c7e1739-c29c-edcb-b541-39f751de902a', 0, '/views/workflow/settings', '/views/workflow/settings', 'Operate.ProcessingSettings', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('de6dea49-731f-2a92-75ce-39f509369c1e', '{}', '20cd053ce02140e48636ed42dd0f7f20', '2020-05-09 09:25:48.584835', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-07-10 17:01:41.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Icon', 'Icon图标', 99, '', 'e4abe794-13f9-c709-0c80-39f5092ce39f', 0, '/icons/index', '/views/icons/index', 'Infrastructure.Icon', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('e2a32f28-8505-4001-8f88-6cda9b63dd77', NULL, NULL, '2022-07-10 11:16:01.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-07-10 12:31:08.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'SystemSettings', '系统设置', 99, '', 'e4abe794-13f9-c709-0c80-39f5092ce39f', 1, 'systemsettings', '', '', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('e4abe794-13f9-c709-0c80-39f5092ce39f', '{}', 'ac66aa78de4246be942491b6efc78bcf', '2020-05-09 09:15:07.552637', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-06-07 23:36:36.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, 'Infrastructure', '基础设置', 2, 'system', '', 1, '/infrastructure', '', '', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('f7352ae7-6875-23e3-afe8-39f789a5d9f5', '{}', '18131777d7024f5abf7f3167201a1522', '2020-09-10 18:04:13.392838', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-10 20:05:23.286402', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-03-19 20:03:13.000000', 'Workflow', '工作流', 3, 'tree', NULL, 1, '/workflow', NULL, NULL, 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');
INSERT INTO `au_menu` VALUES ('fe051ae9-563f-4857-a270-137447f4e030', NULL, NULL, '2022-05-21 19:04:37.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-05-21 21:33:15.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, 'AuditLogAction', '接口日志', 1, '', '942a2853-8f9f-4214-aa7f-53c363947d03', 0, '/auditlog-action/index', '/views/auditlog-action/index', 'AuditLog.AuditLogAction', 0, '0cbba6df-9078-4f4a-9b9c-8d6312d1b184');

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
INSERT INTO `au_permission` VALUES ('02406149-abeb-44af-8770-c91afbd45ca0', NULL, NULL, '2022-03-01 17:02:15.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', '', 'DataPrivilege', '数据授权', 'Operation.DataPrivilege', 5);
INSERT INTO `au_permission` VALUES ('15919b71-1d71-f23a-1759-39f543786b67', '{}', 'e6c082ad5dba4f249e10862c2b4ad30c', '2020-05-20 16:55:36.039229', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-03-19 20:03:01.000000', 'afa5524e-b236-6607-1911-39f543786863', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('17f56317-da3a-57db-066b-39f55c3ebbf0', '{}', 'b610f669b9b4481d8f9ec56b1fc57f6e', '2020-05-25 12:23:05.968785', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-06-28 21:51:57.541222', '344f5ead-8a53-ea44-c25e-39f55c3eb97d', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('1861f3d0-de93-2a8a-4f3a-39f78a0fe93b', '{}', '65948b79f6ff4882b4c17e475ca6cf52', '2020-09-10 20:00:03.899087', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '3081e84b-656b-cbfd-5c8c-39f78a0fe83d', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('19772fe5-9657-4f1a-b81d-561d0ba34a71', NULL, NULL, '2022-05-23 21:41:49.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '0d37c57f-995a-45ce-a8fb-5555c519f5cb', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('2ac09942-6016-1ff6-5b9d-39f53347d853', '{}', '4fe224ebf59745e29974c6b95aba17b2', '2020-05-17 13:28:37.203135', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'FunctionAuthorization', '功能授权', 'Operation.FunctionAuthorization', 2);
INSERT INTO `au_permission` VALUES ('2eaefbcf-650e-25e1-f31a-39f50b43e5c6', '{}', 'b052c3ca19e84bf295dc2af729bfe568', '2020-05-09 18:59:29.862441', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '5c01d798-9d1f-f088-9659-39f50b43e4a3', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('310a5777-ed92-ca37-a876-39f53345c243', '{}', '2bdd45209e1e4b59bd537ece5600d640', '2020-05-17 13:26:20.483446', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '216ad756-bdcc-abd3-c516-39f4ff08b978', 'Operation', 'Create', '新增', 'Operation.Create', 1);
INSERT INTO `au_permission` VALUES ('346f6545-bad6-4273-8fcf-ad304402e8c3', NULL, NULL, '2022-02-12 21:04:05.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, 'ce6c3946-3da7-493a-a5c2-1d5ad4444aee', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('3a9f83cc-0ba6-fda7-9848-39f56645da88', '{}', 'e0b60d0719f94426980312274595c5aa', '2020-05-27 11:07:04.712364', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-27 11:28:31.715106', 'b5047074-9a0c-8721-b2a0-39f56645d7b0', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('3b386f9a-2efe-1f67-2e78-39f533482e9e', '{}', '1b741e0d826d425289134f76feee4faa', '2020-05-17 13:28:59.294866', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'Edit', '编辑', 'Operation.Edit', 3);
INSERT INTO `au_permission` VALUES ('430ec3ec-ba7a-8d72-9346-39f53348fa1b', '{}', '34fcfef25549427b9e8b61022049cb17', '2020-05-17 13:29:51.387740', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'Delete', '删除', 'Operation.Delete', 5);
INSERT INTO `au_permission` VALUES ('4521f526-9e2c-f208-5e53-39f533461812', '{}', '4fcbf073615848faa32b4fc7cc5a14a7', '2020-05-17 13:26:42.450320', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '216ad756-bdcc-abd3-c516-39f4ff08b978', 'Operation', 'Edit', '编辑', 'Operation.Edit', 2);
INSERT INTO `au_permission` VALUES ('4ff5df79-885b-4181-b85a-f9d4b8a2e05c', NULL, NULL, '2022-05-21 19:04:37.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, 'fe051ae9-563f-4857-a270-137447f4e030', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('51041e91-2b99-40b8-b9a4-b4ea04b5bcbd', NULL, NULL, '2022-05-21 15:39:58.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '59fc719d-bbfd-4059-a20c-6672dba08b06', '', 'Edit', '编辑', 'Operation.Edit', 1);
INSERT INTO `au_permission` VALUES ('5d4ed6df-6eb5-fad4-f953-39f53346635f', '{}', '3e1d85a3729842ec960bc6ad08102941', '2020-05-17 13:27:01.727705', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '216ad756-bdcc-abd3-c516-39f4ff08b978', 'Operation', 'Delete', '删除', 'Operation.Delete', 3);
INSERT INTO `au_permission` VALUES ('5ee2d035-47ab-482b-805b-1bf4ffa4d4f6', '{}', 'b052c3ca19e84bf295dc2af729bfe568', '2020-05-09 18:59:29.862441', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '216ad756-bdcc-abd3-c516-39f4ff08b978', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('5f092eaf-a7ca-5cad-8e3a-39f5109e472d', '{}', '9fe9ace282de45259e91a1db397cbdac', '2020-05-10 19:56:19.117957', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('607a2a7f-c1b1-43cf-8a97-7e17a667e86c', NULL, NULL, '2022-06-09 00:28:58.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '0d37c57f-995a-45ce-a8fb-5555c519f5cb', '', 'Delete', '删除', 'Operation.Delete', 5);
INSERT INTO `au_permission` VALUES ('60fa8dbd-2dfb-4b21-b5f8-e321680c9b2c', NULL, NULL, '2022-05-21 15:20:36.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '753ce348-567c-4e19-22f3-39f534336219', '', 'Edit', '编辑', 'Operation.Edit', 2);
INSERT INTO `au_permission` VALUES ('67f82197-0ab1-4d10-a6b9-845911633541', NULL, NULL, '2022-05-05 00:07:14.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '59fc719d-bbfd-4059-a20c-6672dba08b06', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('68e7ac2a-7f02-4ab3-bc8c-7d0736ece563', NULL, NULL, '2022-06-09 00:25:41.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '0d37c57f-995a-45ce-a8fb-5555c519f5cb', '', 'AddDept', '新增部门', 'Operation.AddDept', 3);
INSERT INTO `au_permission` VALUES ('77520662-a275-fa39-195f-39f55c39124e', '{}', 'cd90668e5dce48f8b78b7b37e2f0f861', '2020-05-25 12:16:54.862505', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-03-19 20:03:04.000000', '93352161-b3fb-f99e-5cd9-39f55c390fec', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('79646120-b768-6b63-3724-39f5334426fc', '{}', '3cb693c3f3f44b65940668233bfb1057', '2020-05-17 13:24:35.197675', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '332f6c3f-40cc-2908-99ef-39f4e6b0dfec', 'Operation', 'Create', '新增', 'Operation.Create', 3);
INSERT INTO `au_permission` VALUES ('86c89fbf-e954-42d0-8fd1-b918bb906513', NULL, NULL, '2022-05-21 15:38:13.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '753ce348-567c-4e19-22f3-39f534336219', '', 'EditStatus', '启用/停用', 'Operation.EditStatus', 3);
INSERT INTO `au_permission` VALUES ('8db71675-7626-47ce-a523-827f0edc5dcc', NULL, NULL, '2022-07-10 19:19:59.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '84cfd384-8062-434d-912a-4c39c97a844b', '', 'Create', '新增', 'Operation.Create', 1);
INSERT INTO `au_permission` VALUES ('8eded392-4793-229f-b9a1-39f50936aba8', '{}', 'a88ea79551f641ba960b4c4731aa3c71', '2020-05-09 09:25:48.585026', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, 'de6dea49-731f-2a92-75ce-39f509369c1e', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('94b94ba8-b173-473a-ad64-6c265a16a054', NULL, NULL, '2022-06-08 14:00:58.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '5dd65781-eeca-4e8f-b545-4abdf0e1d8c5', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('99e2461c-d565-4c5b-a741-c81715eb922f', NULL, NULL, '2022-05-21 15:38:39.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '753ce348-567c-4e19-22f3-39f534336219', '', 'Delete', '删除', 'Operation.Delete', 4);
INSERT INTO `au_permission` VALUES ('a6182e8c-ee39-8954-fd0d-39f53348b89c', '{}', '11286c1367794c85bc6e3b6a4e58037d', '2020-05-17 13:29:34.620458', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2022-02-14 16:54:09.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'EditStatus', '启用/禁用', 'Operation.EditStatus', 4);
INSERT INTO `au_permission` VALUES ('ad82e160-208e-f7ca-6b72-39f519df2863', '{}', '59bbf51674e44f69b32c8d4c6ce0a1af', '2020-05-12 15:03:46.019618', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-12 15:04:02.407300', 'de6dea49-731f-2a92-75ce-39f509369c1e', 'View', 'View', '查看', 'View.View', 1);
INSERT INTO `au_permission` VALUES ('ae60fc43-dbc0-3501-9de1-39f53347a8af', '{}', '5efbf3c2438e4d07a0f5374dd3594e98', '2020-05-17 13:28:25.007350', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '1a3e3a43-9b7c-4a21-0c0c-39f5109e463c', 'Operation', 'Create', '新增', 'Operation.Create', 1);
INSERT INTO `au_permission` VALUES ('af543e5e-22de-38a8-6b0d-39f519c854ab', '{}', '5674cba074174f5dbd5f9e2813d26944', '2020-05-12 14:38:50.028212', '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-13 08:45:01.338292', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, NULL, NULL, '332f6c3f-40cc-2908-99ef-39f4e6b0dfec', 'Operation', 'Edit', '编辑', 'Operation.Edit', 1);
INSERT INTO `au_permission` VALUES ('b58a86c9-041b-4db8-af39-5723fc955375', '{}', 'b052c3ca19e84bf295dc2af729bfe568', '2020-05-09 18:59:29.862441', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '332f6c3f-40cc-2908-99ef-39f4e6b0dfec', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('b5cb10e2-f1b8-36b4-679a-39f519e1b9d1', '{}', '1f7c1e4fd0ea4fa6828f98aad498ee0e', '2020-05-12 15:06:34.321069', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '332f6c3f-40cc-2908-99ef-39f4e6b0dfec', 'Operation', 'Delete', '删除', 'Operation.Delete', 2);
INSERT INTO `au_permission` VALUES ('b9d8e5d8-5e0a-4876-9aa4-cb3bc210f716', NULL, NULL, '2022-05-21 19:13:25.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '041d4d04-18db-469f-b44e-8f916747b143', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('c162b6f8-5f65-4e97-a541-799a68923ae3', NULL, NULL, '2022-06-08 11:59:02.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '56410b19-72bb-42c2-a53b-b4082a58bdff', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('c6e6ec1e-99c8-4d78-a0a0-d55113d0a321', NULL, NULL, '2022-06-09 00:25:12.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '0d37c57f-995a-45ce-a8fb-5555c519f5cb', '', 'AddCompany', '新增公司', 'Operation.AddCompany', 2);
INSERT INTO `au_permission` VALUES ('d98eadfd-07ce-142b-2778-39f5343364fd', '{}', 'b869eb86df7e428691dd86b8289a3c56', '2020-05-17 17:45:54.173146', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '753ce348-567c-4e19-22f3-39f534336219', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('dbe722c5-5e04-a1de-ab54-39f4ff08b979', '{}', 'dbd9ad330d224390a5f334c9d7c12751', '2020-05-07 09:59:25.305726', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 0, NULL, NULL, '00000000-0000-0000-0000-000000000000', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('dcfd6495-9c75-4129-8b8f-03367332afa5', NULL, NULL, '2022-07-10 11:20:50.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '84cfd384-8062-434d-912a-4c39c97a844b', '', 'Access', '页面访问', 'View.Access', 1);
INSERT INTO `au_permission` VALUES ('e1bf9c2b-87fd-4a61-a9c3-373ed114f586', NULL, NULL, '2022-03-05 22:55:06.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '753ce348-567c-4e19-22f3-39f534336219', '', 'Create', '新增', 'Operation.Create', 1);
INSERT INTO `au_permission` VALUES ('eab51e5d-4d6e-4361-8ab1-39f751fca927', '{}', '746429c92fa4465c8b2f600a4e85834b', '2020-08-30 22:40:18.215832', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-09-03 21:16:31.844523', 'd4ff0631-df37-280a-2529-39f751fca854', 'View', 'Access', '页面访问', 'View.Access', 0);
INSERT INTO `au_permission` VALUES ('eef2fd0d-ac77-4e11-9971-c8228c249313', NULL, NULL, '2022-05-21 15:39:08.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, '', 0, '', NULL, '753ce348-567c-4e19-22f3-39f534336219', '', 'UpdatePassword', '修改密码', 'Operation.UpdatePassword', 5);
INSERT INTO `au_permission` VALUES ('f18febf0-a698-4222-95d5-16ed69dd6f6b', NULL, NULL, '2022-06-09 00:26:27.000000', '', '2022-06-09 00:28:12.000000', '08d74d8e-4e45-c152-f0b2-22013a71bdde', 0, '', NULL, '0d37c57f-995a-45ce-a8fb-5555c519f5cb', '', 'Modify', '修改', 'Operation.Modify', 4);
INSERT INTO `au_permission` VALUES ('f57df6ef-4e2e-b962-ee82-39f50a959c6a', '{}', 'b10d858b19534a1ba2da9bc7def9dafb', '2020-05-09 15:49:07.818781', '08d74d8e-4e45-c152-f0b2-22013a71bdde', NULL, NULL, 1, '08d74d8e-4e45-c152-f0b2-22013a71bdde', '2020-05-10 18:48:00.953912', 'a6723edb-74e4-3053-c16b-39f50a959be0', 'View', 'Access', '页面访问', 'View.Access', 0);

SET FOREIGN_KEY_CHECKS = 1;
