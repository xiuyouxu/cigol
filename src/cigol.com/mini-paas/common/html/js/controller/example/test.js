define(['app', 'ztree', 'ztreeCheck', 'myDatePicker', "../../service/example/testService"], function (app) {
    app.register.controller('appIndex.example.test', ['$scope', '$state', '$global_var', '$testService', function ($scope, $state, $global_var, $testService) {
        $scope.data=[];
        for(var i=0;i<10;i++) {
            $scope.data[i]={
                'id': i,
                'name': 'name'+i,
                'gender': Math.random() > 0.5 ? 'M' : 'F',
                'comment': 'I am a student...'
            };
        }

        $testService.getHosts().then(function(hosts){
            $scope.hosts=hosts;
        });
    }]);
});