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
    $scope.host = '123456';
    $scope.desc = '123456';
    $scope.method = 'GET';

    $scope.add = function() {
        alert($scope.method);
    };

});
