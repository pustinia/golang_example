module myfirstgo

go 1.20

replace (
    myfirstgo/modules/mylibs => ./modules/mylibs
    myfirstgo/modules/myhttpserver => ./modules/myhttpserver
) 