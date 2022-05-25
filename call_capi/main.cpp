#include <iostream>
#include "libcapi/libcapi.h"

int main() {
    std::cout << "hello" << std::endl;
    auto sum = add(1, 3);
    std::cout << "add(1, 3) = " << sum << std::endl;

    char res[33];
    md5sum("hello world", res);
    std::cout << "md5(\"hello world\") = " << res << std::endl;
    return 0;
}
