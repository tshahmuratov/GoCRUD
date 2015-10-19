angular.module('bookApp',['ui.router','ngResource','bookApp.controllers','bookApp.services']);

angular.module('bookApp').config(function($stateProvider,$httpProvider){
    $stateProvider.state('libraries',{
        url:'/libraries',
        templateUrl:'partials/books.html',
        controller:'LibraryListController'
    }).state('newLibrary',{
        url:'/libraries/new',
        templateUrl:'partials/library-add.html',
        controller:'LibraryCreateController'
    }).state('books',{
        url:'/books',
        templateUrl:'partials/books-list.html',
        controller:'BookListController',
        params: { 
            libId: {
            squash: false,
            },
        }
    }).state('newBook',{
        url:'/books/new',
        templateUrl:'partials/book-add.html',
        controller:'BookCreateController'
    }).state('editBook',{
        url:'/books/:id/edit',
        templateUrl:'partials/book-edit.html',
        controller:'BookEditController'
    });
}).run(function($state){
   $state.go('libraries');
});