#include <iostream>
#include "libcapi/libcapi.h"

int main() {
    std::cout << "hello" << std::endl;
    int sum = add(1, 3);
    std::cout << "add(1, 3) = " << sum << std::endl;

    char res[33];
    md5sum("hello world", res);
    std::cout << "md5sum(\"hello world\") = " << res << std::endl;

    GoComplex64 c64_res = complex64_add(go_complex64(7, 2), go_complex64(4, 5));
    printf("complex64_add(7+2i, 4+5i) = %f+%fi\n", creal(c64_res), cimag(c64_res));

    GoComplex128 c128_res = complex128_add(go_complex128(-1, 3), go_complex128(3, 4));
    printf("complex128_add(-1+3i, 3+4i) = %f+%fi\n", creal(c128_res), cimag(c128_res));

    return 0;
}
