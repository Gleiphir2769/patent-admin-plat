-- 开始初始化数据 ;
INSERT INTO sys_role VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_role VALUES (2, '普通用户', '2', 'user', 0, '', '', 0, '', 0, 0, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

INSERT INTO sys_casbin_rule VALUES (1, 'p', 'user', '/apis/v1/user-agent/*', '*', '', '', '', '', '');
INSERT INTO sys_user VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'admin', '13818888888', 1, '', '', '1', '1@qq.com', '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO sys_user VALUES (2, 'user', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user', '13818888888', 2, '', '', '1', '1@qq.com', '', '2', 0, 0, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'user2', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user2', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'user3', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user3', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'user4', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user4', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);

INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 3, 1, '认领', 1, 0, NULL, NULL);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 5, 1, '认领', 1, 0, NULL, NULL);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 4, 1, '认领', 1, 0, NULL, NULL);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, 11, 1, '认领', 1, 0, NULL, NULL);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 7, 1, '认领', 1, 0, NULL, NULL);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, 10, 1, '认领', 1, 0, NULL, NULL);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, 11, 1, '认领', 1, 0, NULL, NULL);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 1, 1, '关注', 1, 0, NULL, NULL);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 2, 1, '关注', 1, 0, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'good', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'go', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'java', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'c', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'py', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (3, 3, 1, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (4, 4, 1, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (5, 5, 1, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (6, 6, 5, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (7, 7, 5, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (8, 7, 4, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (9, 1, 3, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (10, 2, 1,NULL, NULL);

INSERT INTO package  VALUES (null,'11','11',null,1,0,null,null,null);
INSERT INTO package  VALUES (null,'22','22',null,1,0,null,null,null);
INSERT INTO package  VALUES (null,'inventorgraph','inventorgraph',null,1,0,null,null,null);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '1', '{\"patentId\":1,\"TI\":\"基于1的专利\",\"PNM\":\"1\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of MessageHidden\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author001\",\"CLS\":\"a patent of MessageHidden\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '2', '{\"patentId\":2,\"TI\":\"基于2的专利\",\"PNM\":\"2\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of MemoryFramework\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author002\",\"CLS\":\"a patent of MemoryFramework\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '3', '{\"patentId\":3,\"TI\":\"基于3的专利\",\"PNM\":\"3\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of TwoSoftware\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author003\",\"CLS\":\"a patent of TwoSoftware\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '4', '{\"patentId\":4,\"TI\":\"基于4的专利\",\"PNM\":\"4\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of ThreeSoftware\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author004\",\"CLS\":\"a patent of ThreeSoftware\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '5', '{\"patentId\":5,\"TI\":\"基于5的专利\",\"PNM\":\"5\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author005\",\"CLS\":\"a patent of T0Software\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '6', '{\"patentId\":6,\"TI\":\"基于6的专利\",\"PNM\":\"6\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author006\",\"CLS\":\"a patent of T0Software\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent`  VALUES (7, '7', '{\"patentId\":7,\"TI\":\"基于7的专利\",\"PNM\":\"7\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author007\",\"CLS\":\"a patent of T0Software\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent` VALUES (8, '8', '{\"patentId\":8,\"TI\":\"基于8的专利\",\"PNM\":\"8\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author008\",\"CLS\":\"a patent of T0Software\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent` VALUES (9, '9', '{\"patentId\":9,\"TI\":\"基于9的专利\",\"PNM\":\"9\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"刘思哲",\"CLS\":\"a patent of T0Software\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);
INSERT INTO `patent`  VALUES (10, '10', '{\"patentId\":10,\"TI\":\"基于10的专利\",\"PNM\":\"10\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"刘思哲\",\"CLS\":\"a patent of T0Software\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0, NULL, NULL);



INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (1, 'infringe1', 'important!', 'infringement', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (2, 'infringe2', 'important!', 'infringement', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (3, 'infringe3', 'important!', 'infringement', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (4, 'valuation1', 'important!', 'valuation', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (5, 'valuation2', 'important!', 'valuation', '未审核', 1, 0, '2022-10-18 18:49:24', NULL,NULL);
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `files`) VALUES (6, 'valuation3', 'important!', 'valuation', '未审核', 1, 0, '2022-10-18 18:49:24', NULL, NULL);



-- 数据完成 ;