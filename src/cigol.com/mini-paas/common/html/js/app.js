/**
 * webApp的路由配置
 */
define(['angular', 'router', "angular-ui-bootstrap","ngLocale",'scroll','ngDrag','ngUpload','ngSlider','ngMessages','ngAnimate'], function() {
	var app = angular.module("cigolApp", ['ui.router', 'ui.bootstrap','ngLocale','widget.scrollbar','angular-drag','angularFileUpload','ui.slider','ngMessages','ngAnimate']);
	// app.factory('httpInterceptor', [ '$q', '$injector','$global_var',function($q, $injector,$global_var) {
	// 	var httpInterceptor = {
	// 		'responseError' : function(response) {
	// 			return $q.reject(response);
	// 		},
	// 		'response' : function(response) {
	// 			if(response.config.url.indexOf("restful/cloud")>-1){
	// 				var resObj = angular.fromJson(response.data);
	// 				if(resObj && resObj.result=='timeout'){
	// 					var win = window;
	// 					while (win != win.top){
	// 						win = win.top;
	// 					}
	// 					win.location.href= "/";
	// 				}
	// 			}
	// 			return response;
	// 		},
	// 		'request' : function(config) {
	// 			return config;
	// 		},
	// 		'requestError' : function(config){
	// 			return $q.reject(config);
	// 		}
	// 	};
	// 	return httpInterceptor;
	// }]);
	// app.config([ '$httpProvider', function($httpProvider) {
	// 	$httpProvider.interceptors.push('httpInterceptor');
	// } ]);
	//系统常量设置
	app.constant("$global_var", {
		base_url: '/iaas/',
		user:{}
	});

	//用于注入 controller、filter、directive、service
	app.config(function($controllerProvider, $compileProvider, $filterProvider, $provide) {
		app.register = {
			controller: $controllerProvider.register,
			directive: $compileProvider.directive,
			filter: $filterProvider.register,
			service: $provide.service
		};
	});

	//用于加载ControllerJS
	app.loadControllerJs = function(path) {
		return function($rootScope, $q) {
			var def = $q.defer(),
				deps = [];
			angular.isArray(path) ? (deps = path) : deps.push(path);
			require(deps, function() {
				$rootScope.$apply(function() {
					def.resolve();
				});
			});
			return def.promise;
		}
	};

	//配置前端路由
	app.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
		$urlRouterProvider.when("", "/login");
		$stateProvider
			.state("login", {
				url: "/login",
				controller: 'managementApp.login',
				templateUrl: 'templates/login.html',
				resolve: {
					deps: app.loadControllerJs('js/controller/login')
				}
			})
			.state('appIndex', {
				url: '/appIndex',
				templateUrl: 'templates/appIndex.html',
				controller: 'managementApp.appIndex',
				resolve: {
					deps: app.loadControllerJs('js/controller/appIndex')
				}
			})
			//主页
			.state('appIndex.home', {
				url: '/home',
				templateUrl: 'templates/home/home.html',
				controller: 'appIndex.home',
				resolve: {
					deps: app.loadControllerJs('js/controller/home/home')
				}
			})
			// example submodule
			.state('appIndex.example', {
                url: '/example',
                templateUrl: 'templates/example/example.html',
                controller: 'appIndex.example',
                resolve: {
                    deps: app.loadControllerJs('js/controller/example/example')
                }
            })
			// example test
			.state('appIndex.example.test', {
				url: '/test',
				templateUrl: 'templates/example/test.html',
				controller: 'appIndex.example.test',
				resolve: {
					deps: app.loadControllerJs('js/controller/example/test')
				}
			})
			// example test
			.state('appIndex.example.test2', {
				url: '/test',
				templateUrl: 'templates/example/test.html',
				controller: 'appIndex.example.test',
				resolve: {
					deps: app.loadControllerJs('js/controller/example/test')
				}
			})
			// example test
			.state('appIndex.example.test3', {
				url: '/test',
				templateUrl: 'templates/example/test.html',
				controller: 'appIndex.example.test',
				resolve: {
					deps: app.loadControllerJs('js/controller/example/test')
				}
			})
	}]);


	return app;



});

