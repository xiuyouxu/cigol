define(['app', 'js/service/loginService'], function(app) {
	app.register.controller('managementApp.login', ['$scope', '$state', '$loginService', '$global_var', function($scope, $state, $loginService, $global_var) {
		'use strict';
		$scope.username = "";
		$scope.password = "";
		$scope.message = "";
		$scope.userNameMessage = "";
		$scope.passwordMessage = "";

		$scope.checkUserName = function (event) {
			if($scope.username == ""){
				$scope.userNameMessage="请填写登陆名称！";
			}
		}
		$scope.checkPassword = function (event) {
			if($scope.password == ""){
				$scope.passwordMessage="请填写登陆密码！";
			}
		}
		$scope.usernameOnFocus = function (event) {
			$scope.userNameMessage = "";
		}
		$scope.passwordOnFocus = function (event) {
			event.type = "password";
			$scope.passwordMessage = "";
		}
		$scope.submitForm = function(valid) {
			if(valid) {
				if($scope.username == ""){
					$scope.userNameMessage="请填写登陆名称！";
					return;
				}
				if($scope.password == ""){
					$scope.passwordMessage="请填写登陆密码！";
					return;
				}
				var postData = {
					"loginName": $scope.username,
					"password": $scope.password
				}

				$state.go('appIndex.example.test')

				// $loginService.loginAuth(postData).then(function(result) {
					// var result=$loginService.loginAuth(postData);
					// if(result.status === "success") {
						// $loginService.getMenu().then(function(menu) {
							// var menu=$loginService.getMenu();
       //                      console.dir("菜单数据:"+menu)
							// for(var item in menu){
							// 	console.dir(item);
							// 	console.dir(menu[item]);
							// }
       //                      $global_var.user = JSON.parse(result.content);
       //                      $global_var.user.permission = dataProccess(angular.fromJson(menu.content), "");
       //                      $global_var.user.isLogged = true;
       //                      localStorage.user = JSON.stringify($global_var.user)
       //                      $state.go('appIndex.home')

						// })

					// }else{
     //                    $scope.message = result.message;
					// }
				// })
			}
		}

		var dataProccess = function(menu, permission) {
			var result = {
					"mainMenu": [],
					"menu": {
						"resource": [],
						"dg": [],
						"dataManage": [],
                        "etl": [],
                        "taskManage": [],
						"console": [],
						"dataMart": [],
						"personalCenter": [],
						"serviceCenter": [],
						"subjectAnalysis": []
					}
				}
				//遍历全部基础菜单数据
			var parentId = []
			var secondLevelMenu = []
			var thirdLevelMenu = []
			var user_type = $global_var.user.type;
			for(var i = 0; i < menu.length; i++) {
				//提取出其中的父顶级菜单
				if(menu[i].level === '0') {
					var obj = {
						name: menu[i].name,
						iconUrl: menu[i].img,
						router: menu[i].url,
						id: menu[i].id,
						index:menu[i].menuIndex,
						pId: menu[i].pId
					}
					if(user_type==='0'){
						result.mainMenu.push(obj);
						parentId.push(menu[i].pId);
					}else if(user_type=='2' && menu[i].id=='6'){
						result.mainMenu.push(obj);
						parentId.push(menu[i].pId);
					}else if(user_type=='1' && menu[i].id!='6'){
						result.mainMenu.push(obj);
						parentId.push(menu[i].pId);
					}
				}

				//二级菜单
				var parentIdLen = menu[i].level;
				if(parentIdLen==1 && menu[i].pId != 0){
					var child_obj = {
						name: menu[i].name,
						iconUrl: menu[i].img,
						router: menu[i].url,
						id: menu[i].id,
						index:menu[i].menuIndex,
						pId: menu[i].pId,
						child:[]
					}
					if(user_type==='0' && menu[i].id!='61'){
						secondLevelMenu.push(child_obj);
						parentId.push(menu[i].pId)
					}else if(user_type=='2' && (menu[i].id=='61' || menu[i].id=='64')){
						secondLevelMenu.push(child_obj);
						parentId.push(menu[i].pId)
					}else if(user_type=='1'){
						secondLevelMenu.push(child_obj);
						parentId.push(menu[i].pId)
					}

				}

				//三级菜单
				if(parentIdLen==2){
					var child_obj = {
						name: menu[i].name,
						iconUrl: menu[i].img,
						router: menu[i].url,
						id: menu[i].id,
						index:menu[i].menuIndex,
						pId: menu[i].pId,
						child:[]
					}
					thirdLevelMenu.push(child_obj);
				}
			}
			//组合三级菜单和二级菜单
			for(var var1 =0;var1<secondLevelMenu.length;var1++){
				for(var var2 =0;var2<thirdLevelMenu.length;var2++ ){
					if(thirdLevelMenu[var2].pId == secondLevelMenu[var1].id){
						secondLevelMenu[var1].child.push(thirdLevelMenu[var2]);
					}
				}
				if($.inArray("resource", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.resource.push(secondLevelMenu[var1])
				} else if($.inArray("dg", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.dg.push(secondLevelMenu[var1])
				} else if($.inArray("dataManage", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.dataManage.push(secondLevelMenu[var1])
				} else if($.inArray("etl", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.etl.push(secondLevelMenu[var1])
				} else if($.inArray("taskManage", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.taskManage.push(secondLevelMenu[var1])
				} else if($.inArray("console", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.console.push(secondLevelMenu[var1])
				}
				else if($.inArray("dataMart", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.dataMart.push(secondLevelMenu[var1])
				}
				else if($.inArray("personalCenter", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.personalCenter.push(secondLevelMenu[var1])
				}
				else if($.inArray("serviceCenter", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.serviceCenter.push(secondLevelMenu[var1])
				}
				else if($.inArray("subjectAnalysis", secondLevelMenu[var1].router.split(".")) !== -1) {
					result.menu.subjectAnalysis.push(secondLevelMenu[var1])
				}
			}
			return result
		}
	}]);
});
