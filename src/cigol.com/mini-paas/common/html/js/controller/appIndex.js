define(['app', 'jquery', 'directive','datasourceDirective','fileDatasourceDirective', 'filter','cigoldialog', "js/service/appIndexService"], function(app, $) {
	app.register.controller('managementApp.appIndex', ['$scope', '$global_var', '$state', '$appIndexService', function($scope, $global_var, $state, $appIndexService) {
		'use strict';		
		//换肤
		$scope.changecss = function(color) {
			var baseUrl = "css/theme/"
			var url = baseUrl + color + ".css"
			$("#skin").prop("href", url)
			localStorage.user_skin = color
		}
		if(!localStorage.user) {
			$global_var.user.isLogged = false;
			$state.go('login')
		} else {
			$global_var.user = JSON.parse(localStorage.user);
			if($global_var.user){
				$scope.userName = $global_var.user.userName
				$scope.permission = $global_var.user.permission
				if(localStorage.user_skin){
					$scope.changecss(localStorage.user_skin)
				}
			}

		}
		//退出登录
		$scope.back = function() {
			$global_var.user = null;
			localStorage.user = null;
			removeCookies("user_id");
			removeCookies("user_role_id");
			removeCookies("username");
			removeCookies("loginname");
			$state.go('login')
		}
		var removeCookies = function (name) {
				var exp = new Date();
				exp.setTime(exp.getTime() - 1);
				var cval=getCookie(name);
				if(cval!=null)
					document.cookie= name + "="+cval+";expires="+exp.toGMTString();

		}
		var getCookie = function (name)
		{
			var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
			if(arr=document.cookie.match(reg))
				return unescape(arr[2]);
			else
				return null;
		}
		$scope.global_function = {
			checkField: function(sorArr, sorFld, sorVal, tarFld) {
				var out;
				for(var i = 0; i < sorArr.length; i++) {
					if(sorArr[i][sorFld] === sorVal) {
						out = sorArr[i][tarFld]
					}
				}
				return out
			},
			massdelete: function(arr) {
				var str = ""
				for(var i = 0; i < arr.length; i++) {
					str += (arr[i] + ',')
				}
				str = str.slice(0, str.length - 1)
				return str
			}
		}

		//根据id查询用户所有信息
		$scope.userEdit = function() {
			$appIndexService.queryUserById($global_var.user.id).then(function(result) {
				console.group("【用户信息查询】'base/user/byId'");
				console.group("提交数据");
				console.dir($scope.userId);
				console.groupEnd();
				console.group("接收数据");
				console.dir(result);
				console.groupEnd();
				console.groupEnd();
				$scope.dataItems = result.content;
				$scope.userNameInfoEdit = $scope.dataItems.user;
				$scope.pwdInfoEdit = "000000";
				$scope.nameInfoEdit = $scope.dataItems.name;
				$scope.mobilePhoneInfoEdit = $scope.dataItems.mobilePhone;
				$scope.unitInfoEdit = $scope.dataItems.unit;

			}, function() {
				$scope.errMsg = "系统繁忙，请稍后再试!";
			});
		};

		$scope.userInfoReset = function() {
			if($scope.dataItems){
				$scope.userNameInfoEdit = $scope.dataItems.user;
				$scope.pwdInfoEdit = "000000";
				$scope.nameInfoEdit = $scope.dataItems.name;
				$scope.mobilePhoneInfoEdit = $scope.dataItems.mobilePhone;
				$scope.unitInfoEdit = $scope.dataItems.unit;
			}
		}

		//--------------------编辑向后台提交数据--------------------//
		$scope.userInfoEdit = function(valid) {
			if(valid) {
				if($scope.pwdInfoEdit === "000000") {
					var userEditItem = {
						"id": $scope.userId,
						"user": $scope.userNameInfoEdit,
						"name": $scope.nameInfoEdit,
						"mobilePhone": $scope.mobilePhoneInfoEdit,
						"unit": $scope.unitInfoEdit

					};
				} else {
					var userEditItem = {
						"id": $scope.userId,
						"user": $scope.userNameInfoEdit,
						"pwd": $scope.pwdInfoEdit,
						"name": $scope.nameInfoEdit,
						"mobilePhone": $scope.mobilePhoneInfoEdit,
						"unit": $scope.unitInfoEdit

					};
				}

				$appIndexService.edit(userEditItem).then(function(result) {
					console.group("【用户信息修改】'base/user/updating'");
					console.group("提交数据");
					console.dir(userEditItem);
					console.groupEnd();
					console.group("接收数据");
					console.dir(result);
					console.groupEnd();
					console.groupEnd();

					if(result.content === "true") {} else {}
					$('#globleEditInfo').modal('hide');
				});
			}

		};

	}]);
});