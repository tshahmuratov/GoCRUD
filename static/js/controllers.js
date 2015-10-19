angular.module('bookApp.controllers',[]).controller('LibraryListController',function($scope,$state,popupService,$window,Book, Library){

    $scope.libraries=Library.query();
    $scope.deleteLibrary=function(library){
        if(popupService.showPopup('Действительно хотите удалить?')){
            library.$delete(function(){
                $window.location.href='';
            });
        }
    }
    $scope.deleteBook=function(book){
        if(popupService.showPopup('Действительно хотите удалить?')){
            book.$delete(function(){
                $window.location.href='';
            });
        }
    }
    $scope.viewLibrary=function(library) {
        $scope.books = Library.get({id:library.id.Id});
    }

}).controller('SelectLibraryController',function($scope, $stateParams, Book, Library, BooksInLibrary){

    $scope.library=new Library();

    $scope.addLibrary=function(){
        $scope.library.$save(function(){
            $state.go('libraries');
        });
    }

}).controller('LibraryCreateController',function($scope,$stateParams,Book){

    $scope.library=new Library();

    $scope.addLibrary=function(){
        $scope.library.$save(function(){
            $state.go('libraries');
        });
    }

}).controller('BookListController',function($scope,$state,$stateParams,Book, Library){

    $scope.books=Book.query();

    $scope.deleteBook=function(book){
        if(popupService.showPopup('Really delete this?')){
            book.$delete(function(){
                $window.location.href='';
            });
        }
    }
    $scope.setSelected = function(id) {
        $scope.bookValue = id;
    }
    $scope.addBookToLibrary=function() {
        $.ajax({
            url: "/api/v1/libraries/"+$stateParams.libId+"/"+$scope.bookValue,
            method: 'POST',
            data: '{}'
        }).done(function() {
            $state.go('libraries');
        });
    };

}).
controller('BookCreateController',function($scope,$state,$stateParams,Book){

    $scope.book=new Book();

    $scope.addBook=function(){
        $scope.book.$save(function(){
            $state.go('books');
        });
    }

}).controller('BookEditController',function($scope,$state,$stateParams,Book){

    $scope.updateBook=function(){
        $scope.book.$update(function(){
            $state.go('books');
        });
    };

    $scope.loadBook=function(){
        $scope.book=Book.get({id:$stateParams.id});
    };

    $scope.loadBook();
}).controller('addBookToLibraryController',function($scope,$state,$stateParams,Book){

    $scope.updateBook=function(){
        $scope.book.$update(function(){
            $state.go('books');
        });
    };

    $scope.loadBook=function(){
        $scope.book=Book.get({id:$stateParams.id});
    };
    $scope.loadBook();
}).controller('SelectLibraryController', ['$scope', function($scope) {
   $scope.data = {
    option1: 'Выберите библиотеку',
   };
}]);