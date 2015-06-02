var routerApp = angular.module('httpTest', ['ngRoute']);

routerApp.config(['$routeProvider', '$locationProvider',
    function($routeProvider, $locationProvider) {
        $locationProvider.html5Mode({
            enabled: true,
            requireBase: false
        });

        $routeProvider
            .when('/', {
                templateUrl: '/views/add.html',
                controller: 'appCtrl'
            })
            .when('/list', {
                templateUrl: '/views/list.html',
                controller: 'appCtrl'
            })
            .when('/about', {
                templateUrl: '/views/about.html',
                controller: 'appCtrl'
            })
            .otherwise({
                redirectTo: '/'
            });
    }
]);

routerApp.controller('appCtrl', function($scope, $http) {
    $scope.host = '';
    $scope.desc = '';
    $scope.message = '';
    $scope.method = 0;
    $scope.params = [{}];


    $scope.add = function() {
        for (var i = 0; i < $scope.params.length; i++) {
            console.log($scope.params[i].key);
        }
        $http.post('/api/add', {
            host: $scope.host,
            desc: $scope.desc,
            method: parseInt($scope.method),
            param: $scope.params,
        }).success(function(data) {
            $scope.message = data.message;
            if (data.code == 200) {
            }
        }).error(function() {
            $scope.message = '请求错误!';
        });
    };

    $scope.keyvalueClick = function() {
        var last = $scope.params[$scope.params.length - 1];
        if (last.key != null || last.value != null) {
            $scope.params.push({})
        }
    };

    $scope.delParam = function() {
        var index = this.$index;
        if (~index) $scope.params.splice(index, 1);
    };

});
