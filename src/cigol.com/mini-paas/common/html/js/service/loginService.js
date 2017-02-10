define(['app'],function(app){
    app.register.service('$loginService', ['$http','$q','$global_var',function($http,$q,$global_var){
    	this.loginAuth = function(data){
			// var deferred=$q.defer();
			// $http({
			// 	method: 'POST',
			// 	url: $global_var.base_url + "base/userRestServer/loginCheck",
			// 	data:data
			// }).success(function(result) {
			// 		deferred.resolve(result);
			// }).error(function(err){
			// 	deferred.reject(false);
			// 	console.log(err);
			// });
			// return deferred.promise;

			// response data should be like
			return {"content":"{\"address\":\"pingtai\",\"createDate\":1482112988000,\"email\":\"pingtai@qq.com\",\"id\":\"9A7B9CB3983344FE8C88286A6C188BF4\",\"loginName\":\"pingtai\",\"passWord\":\"03142410d7285f00e4363e005783c83a\",\"phone\":\"18146785609\",\"roleId\":\"AB9F6BDADBF24751A91AE9C40EBEEAC7\",\"tenantId\":\"3D002E13B106450ABFD9FE73C23D460A\",\"type\":\"0\",\"userName\":\"pingtai\"}","message":"用户信息正确！","status":"success"};
		};
		
		this.getMenu = function(){
			// var deferred=$q.defer();
			// $http({
			// 	method: 'GET',
			// 	url: $global_var.base_url + "base/menuRestServer/getMenu",
			// }).success(function(result) {
			// 		deferred.resolve(result);
			// }).error(function(err){
			// 	deferred.reject(false);
			// 	console.log(err);
			// });
			// return deferred.promise;

			// response data should be like
			return {"content":"[{\"createDate\":1484641173000,\"id\":\"333\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"æ•°æ®æµè§ˆ\",\"pId\":\"33\",\"url\":\"appIndex.dataManage.fileStorageDataBrowse\"},{\"createDate\":1484641169000,\"id\":\"111\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"æŒ‡æ ‡é¡¹ç®¡ç†\",\"pId\":\"11\",\"url\":\"appIndex.resource.itemManager\"},{\"createDate\":1484641174000,\"id\":\"364\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"æ•°æ®æµè§ˆ\",\"pId\":\"36\",\"url\":\"appIndex.dataManage.indexDataBrowseManager\"},{\"createDate\":1484641171000,\"id\":\"131\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"ä¸»é¢˜åº“ç›®å½•\",\"pId\":\"13\",\"url\":\"appIndex.resource.topicDataDirectory\"},{\"createDate\":1484641173000,\"id\":\"341\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"å­˜å‚¨é…ç½®\",\"pId\":\"34\",\"url\":\"appIndex.dataManage.databaseConfig\"},{\"createDate\":1484641172000,\"id\":\"1\",\"img\":\"img/index/2.png\",\"level\":\"0\",\"menuIndex\":1,\"name\":\"èµ„æºç®¡ç†\",\"pId\":\"0\",\"url\":\"appIndex.resource\"},{\"createDate\":1484641174000,\"id\":\"3\",\"img\":\"img/index/4.png\",\"level\":\"0\",\"menuIndex\":1,\"name\":\"æ•°æ®ç®¡ç†\",\"pId\":\"0\",\"url\":\"appIndex.dataManage.serverComfing\"},{\"createDate\":1484641178000,\"id\":\"01B75F7135074CB79C12C7EFAD02EF77\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"æ•°æ®ç”³è¯·\",\"pId\":\"D1C9BAE7E3FC40E98B2A8E966E8D6FFB\",\"url\":\"appIndex.personalCenter.themeDataApplication\"},{\"createDate\":1484641173000,\"id\":\"343\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"æ•°æ®æµè§ˆ\",\"pId\":\"34\",\"url\":\"appIndex.dataManage.databaseDataBrowse\"},{\"createDate\":1484641172000,\"id\":\"21\",\"img\":\"img/index/menu/dg/4.png\",\"level\":\"1\",\"menuIndex\":1,\"name\":\"æ•°æ®æºç®¡ç†\",\"pId\":\"2\",\"url\":\"appIndex.dg.datasource\"},{\"createDate\":1484641178000,\"id\":\"CE976C2B13434F5E86EA4B7673193E90\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"åŽŸå§‹æ•°æ®\",\"pId\":\"095E0E04913C4D4DB4A1BF102BC41BD6\",\"url\":\"appIndex.personalCenter.primaryData\"},{\"createDate\":1484641179000,\"id\":\"43C53289A71348AA87465FDFDBD0E387\",\"img\":\"img/index/7.png\",\"level\":\"0\",\"menuIndex\":1,\"name\":\"ä¸ªäººä¸­å¿ƒ\",\"pId\":\"0\",\"url\":\"appIndex.personalCenter\"},{\"createDate\":1484641174000,\"id\":\"361\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"ç´¢å¼•åº“ç®¡ç†\",\"pId\":\"36\",\"url\":\"appIndex.dataManage.indexDataBaseManage\"},{\"createDate\":1484641173000,\"id\":\"332\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"è·¯å¾„é…ç½®\",\"pId\":\"33\",\"url\":\"appIndex.dataManage.fileStoragePathConfig\"},{\"createDate\":1484641172000,\"id\":\"2\",\"img\":\"img/index/3.png\",\"level\":\"0\",\"menuIndex\":1,\"name\":\"æ•°æ®é‡‡é›†\",\"pId\":\"0\",\"url\":\"appIndex.dg\"},{\"createDate\":1484641169000,\"id\":\"11\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"1\",\"menuIndex\":1,\"name\":\"æ•°æ®é‡‡é›†ç®¡ç†\",\"pId\":\"1\",\"url\":\"appIndex.resource.dgManager\"},{\"createDate\":1484641178000,\"id\":\"5C35B194B4704246B85D23BBBDB6B5F2\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"æ•°æ®ä¸ŠæŠ¥\",\"pId\":\"AAD545749B3944A1BD6BD3D55CBBDC9E\",\"url\":\"appIndex.personalCenter.dataReport\"},{\"createDate\":1484641174000,\"id\":\"363\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"ç´¢å¼•è§„åˆ™å®šåˆ¶\",\"pId\":\"36\",\"url\":\"appIndex.dataManage.indexRuleConfig\"},{\"createDate\":1484641178000,\"id\":\"9F2F60C0DB90491FB90F49A85D5E955E\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"å…±äº«ç”³è¯·\",\"pId\":\"D1C9BAE7E3FC40E98B2A8E966E8D6FFB\",\"url\":\"appIndex.personalCenter.shareDataApplication\"},{\"createDate\":1484641176000,\"id\":\"6\",\"img\":\"img/index/7.png\",\"level\":\"0\",\"menuIndex\":1,\"name\":\"å¹³å°ç®¡ç†\",\"pId\":\"0\",\"url\":\"appIndex.console.permissionManagement\"},{\"createDate\":1484641177000,\"id\":\"614BE677218743BDAB806892B4A9B717\",\"img\":\"img/index/7.png\",\"level\":\"0\",\"menuIndex\":1,\"name\":\"æ•°æ®é›†å¸‚\",\"pId\":\"0\",\"url\":\"appIndex.dataMart\"},{\"createDate\":1484641173000,\"id\":\"342\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"è¡¨ç®¡ç†\",\"pId\":\"34\",\"url\":\"appIndex.dataManage.databaseTableManager\"},{\"createDate\":1484641178000,\"id\":\"095E0E04913C4D4DB4A1BF102BC41BD6\",\"img\":\"img/index/menu/console/1.png\",\"level\":\"1\",\"menuIndex\":1,\"name\":\"æˆ‘çš„èµ„æº\",\"pId\":\"43C53289A71348AA87465FDFDBD0E387\",\"url\":\"appIndex.personalCenter.myResources\"},{\"createDate\":1484641176000,\"id\":\"76FA6564F50E41A3B78CB6D101FE6D70\",\"img\":\"img/index/menu/console/1.png\",\"level\":\"1\",\"menuIndex\":1,\"name\":\"åŽŸå§‹æ•°æ®\",\"pId\":\"614BE677218743BDAB806892B4A9B717\",\"url\":\"appIndex.dataMart.primaryData\"},{\"createDate\":1484641172000,\"id\":\"151\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"ç”³è¯·å®¡æ‰¹\",\"pId\":\"15\",\"url\":\"appIndex.resource.applicationApprove\"},{\"createDate\":1484641173000,\"id\":\"331\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"å­˜å‚¨é…ç½®\",\"pId\":\"33\",\"url\":\"appIndex.dataManage.fileStorageConfig\"},{\"createDate\":1484641178000,\"id\":\"97498B19E3314B508EFB273C9C3F4B13\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"ä¸»é¢˜æ•°æ®\",\"pId\":\"095E0E04913C4D4DB4A1BF102BC41BD6\",\"url\":\"appIndex.personalCenter.themeData\"},{\"createDate\":1484641169000,\"id\":\"121\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"æ¸…æ´—è§„åˆ™\",\"pId\":\"12\",\"url\":\"appIndex.resource.cleanRule\"},{\"createDate\":1484641174000,\"id\":\"362\",\"img\":\"\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"ç´¢å¼•è¡¨ç®¡ç†\",\"pId\":\"36\",\"url\":\"appIndex.dataManage.indexTableManage\"},{\"createDate\":1484641171000,\"id\":\"141\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":1,\"name\":\"æ•°æ®å‘å¸ƒ\",\"pId\":\"14\",\"url\":\"appIndex.resource.dataPublish\"},{\"createDate\":1484641178000,\"id\":\"D1C9BAE7E3FC40E98B2A8E966E8D6FFB\",\"img\":\"img/index/menu/console/2.png\",\"level\":\"1\",\"menuIndex\":2,\"name\":\"æˆ‘çš„ç”³è¯·\",\"pId\":\"43C53289A71348AA87465FDFDBD0E387\",\"url\":\"appIndex.personalCenter.myApplication\"},{\"createDate\":1484641169000,\"id\":\"112\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":2,\"name\":\"æ•°æ®æµè§ˆ\",\"pId\":\"11\",\"url\":\"appIndex.resource.itemDataBrowse\"},{\"createDate\":1484641176000,\"id\":\"62\",\"img\":\"img/index/menu/console/2.png\",\"level\":\"1\",\"menuIndex\":2,\"name\":\"è§’è‰²ç®¡ç†\",\"pId\":\"6\",\"url\":\"appIndex.console.roleManagement\"},{\"createDate\":1484641176000,\"id\":\"59312834D8D64551B8D4A11B507078DC\",\"img\":\"img/index/menu/console/2.png\",\"level\":\"1\",\"menuIndex\":2,\"name\":\"ä¸»é¢˜æ•°æ®\",\"pId\":\"614BE677218743BDAB806892B4A9B717\",\"url\":\"appIndex.dataMart.themeData\"},{\"createDate\":1484641172000,\"id\":\"152\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":2,\"name\":\"å…±äº«å®¡æ‰¹\",\"pId\":\"15\",\"url\":\"appIndex.resource.shareApprove\"},{\"createDate\":1484641170000,\"id\":\"122\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":2,\"name\":\"æ¸…æ´—ç»“æžœ\",\"pId\":\"12\",\"url\":\"appIndex.resource.cleanResult\"},{\"createDate\":1484641171000,\"id\":\"13\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"1\",\"menuIndex\":2,\"name\":\"æ•°æ®ç®¡ç†ä¸­å¿ƒ\",\"pId\":\"1\",\"url\":\"appIndex.resource.dataManager\"},{\"createDate\":1484641171000,\"id\":\"132\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":2,\"name\":\"æ•°æ®æµè§ˆ\",\"pId\":\"13\",\"url\":\"appIndex.resource.topicDataBrowse\"},{\"createDate\":1484641172000,\"id\":\"23\",\"img\":\"img/index/menu/dg/2.png\",\"level\":\"1\",\"menuIndex\":2,\"name\":\"æŒ‡æ ‡é¡¹ç®¡ç†\",\"pId\":\"2\",\"url\":\"appIndex.dg.item\"},{\"createDate\":1484641173000,\"id\":\"32\",\"img\":\"img/index/menu/dataManage/8.png\",\"level\":\"1\",\"menuIndex\":2,\"name\":\"æœåŠ¡å™¨ä¿¡æ¯ç®¡ç†\",\"pId\":\"3\",\"url\":\"appIndex.dataManage.serverComfing\"},{\"createDate\":1484641172000,\"id\":\"22\",\"img\":\"img/index/menu/dg/1.png\",\"level\":\"1\",\"menuIndex\":3,\"name\":\"èŠ‚ç‚¹ç®¡ç†\",\"pId\":\"2\",\"url\":\"appIndex.dg.cluster\"},{\"createDate\":1484641173000,\"id\":\"33\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"1\",\"menuIndex\":3,\"name\":\"æ–‡ä»¶ç®¡ç†\",\"pId\":\"3\",\"url\":\"appIndex.dataManage.fileStorage\"},{\"createDate\":1484641171000,\"id\":\"12\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"1\",\"menuIndex\":3,\"name\":\"æ•°æ®èžåˆç®¡ç†\",\"pId\":\"1\",\"url\":\"appIndex.resource.etlManager\"},{\"createDate\":1484641176000,\"id\":\"63\",\"img\":\"img/index/menu/console/1.png\",\"level\":\"1\",\"menuIndex\":3,\"name\":\"ç”¨æˆ·ç®¡ç†\",\"pId\":\"6\",\"url\":\"appIndex.console.deptManager\"},{\"createDate\":1484641170000,\"id\":\"123\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":3,\"name\":\"å¯¹æ¯”è§„åˆ™\",\"pId\":\"12\",\"url\":\"appIndex.resource.compareRule\"},{\"createDate\":1484641179000,\"id\":\"AAD545749B3944A1BD6BD3D55CBBDC9E\",\"img\":\"img/index/menu/console/2.png\",\"level\":\"1\",\"menuIndex\":3,\"name\":\"æˆ‘çš„å…±äº«\",\"pId\":\"43C53289A71348AA87465FDFDBD0E387\",\"url\":\"appIndex.personalCenter.myShare\"},{\"createDate\":1484641171000,\"id\":\"124\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"2\",\"menuIndex\":4,\"name\":\"å¯¹æ¯”ç»“æžœ\",\"pId\":\"12\",\"url\":\"appIndex.resource.compareResult\"},{\"createDate\":1484641176000,\"id\":\"64\",\"img\":\"img/index/menu/console/4.png\",\"level\":\"1\",\"menuIndex\":4,\"name\":\"èœå•æŽˆæƒ\",\"pId\":\"6\",\"url\":\"appIndex.console.permissionManagement\"},{\"createDate\":1484641171000,\"id\":\"14\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"1\",\"menuIndex\":4,\"name\":\"èµ„æºå‘å¸ƒç®¡ç†\",\"pId\":\"1\",\"url\":\"appIndex.resource.publishManager\"},{\"createDate\":1484641173000,\"id\":\"34\",\"img\":\"img/index/menu/dataManage/3.png\",\"level\":\"1\",\"menuIndex\":4,\"name\":\"æ•°æ®åº“ç®¡ç†\",\"pId\":\"3\",\"url\":\"appIndex.dataManage.database\"},{\"createDate\":1484641172000,\"id\":\"24\",\"img\":\"img/index/menu/dg/3.png\",\"level\":\"1\",\"menuIndex\":4,\"name\":\"æµç¨‹ç®¡ç†\",\"pId\":\"2\",\"url\":\"appIndex.dg.workflow\"},{\"createDate\":1484641172000,\"id\":\"25\",\"img\":\"img/index/menu/dg/5.png\",\"level\":\"1\",\"menuIndex\":5,\"name\":\"æµç¨‹æ—¥å¿—\",\"pId\":\"2\",\"url\":\"appIndex.dg.logs\"},{\"createDate\":1484641172000,\"id\":\"15\",\"img\":\"img/index/menu/dataManage/2.png\",\"level\":\"1\",\"menuIndex\":5,\"name\":\"å®¡æ‰¹ä¸­å¿ƒ\",\"pId\":\"1\",\"url\":\"appIndex.resource.approveManager\"},{\"createDate\":1484641176000,\"id\":\"65\",\"img\":\"img/index/menu/console/5.png\",\"level\":\"1\",\"menuIndex\":5,\"name\":\"æœåŠ¡å™¨ç®¡ç†\",\"pId\":\"6\",\"url\":\"appIndex.console.serverManager\"},{\"createDate\":1484641176000,\"id\":\"66\",\"img\":\"img/index/menu/console/6.png\",\"level\":\"1\",\"menuIndex\":6,\"name\":\"èµ„æºåŒ…ç®¡ç†\",\"pId\":\"6\",\"url\":\"appIndex.console.resourceManager\"},{\"createDate\":1484641174000,\"id\":\"36\",\"img\":\"img/index/menu/dataManage/7.png\",\"level\":\"1\",\"menuIndex\":6,\"name\":\"ç´¢å¼•ç®¡ç†\",\"pId\":\"3\",\"url\":\"appIndex.dataManage.tasks\"},{\"createDate\":1484641176000,\"id\":\"67\",\"img\":\"img/index/menu/console/7.png\",\"level\":\"1\",\"menuIndex\":7,\"name\":\"éƒ¨ç½²é…ç½®\",\"pId\":\"6\",\"url\":\"appIndex.console.deployConfigManager\"},{\"createDate\":1484641176000,\"id\":\"68\",\"img\":\"img/index/menu/console/9.png\",\"level\":\"1\",\"menuIndex\":8,\"name\":\"æœåŠ¡ç®¡ç†\",\"pId\":\"6\",\"url\":\"appIndex.console.serviceManager\"}]","message":"","status":"success"};
		};

		this.getPermission = function(roleId){
			var deferred=$q.defer();
			$http({
				method: 'GET',
				url: $global_var.base_url + "base/role/" + roleId
			}).success(function(result) {
					deferred.resolve(result);
			}).error(function(err){
				deferred.reject(false);
				console.log(err);
			});
			return deferred.promise;
		};

		

    }]);
});

