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
                controller: 'add'
            })
            .when('/list', {
                templateUrl: '/views/list.html',
                controller: 'list'
            })
            .when('/about', {
                templateUrl: '/views/about.html',
                controller: 'add'
            })
            .otherwise({
                redirectTo: '/'
            });
    }
]);

routerApp.controller('add', function($scope, $http, $location) {
    $scope.host = '';
    $scope.desc = '';
    $scope.message = '';
    $scope.method = 0;
    $scope.params = [{}];

    $scope.isActive = function (viewLocation) { 
        return viewLocation === $location.path();
    };

    $scope.add = function() {
        $http.post('/api/add', {
            host: $scope.host,
            desc: $scope.desc,
            method: parseInt($scope.method),
            param: $scope.params,
        }).success(function(data) {
            $scope.message = data.message;
            if (data.Code == 200) {
                $scope.host = '';
                $scope.desc = '';
                $scope.message = '';
                $scope.method = 0;
                $scope.params = [{}];
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

}).controller('list', function($scope, $http, $location) {
    
    $scope.isActive = function (viewLocation) { 
        return viewLocation === $location.path();
    };

    $scope.init = function() {
        $http.get('/api/list').success(function(data) {
            if (data.Code == 200) {
                $scope.requests = data.Data;
            }
        }).error(function() {
            $scope.message = '请求错误!';
        });
    };

    $scope.init();

});
