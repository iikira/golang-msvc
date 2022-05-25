go build -buildmode=c-archive -ldflags "-s -w" -o ./call_capi/libcapi/libcapi.a ./capi

gcc ./capi/libcapi.def ./call_capi/libcapi/libcapi.a -shared -lwinmm -lWs2_32 -o ./call_capi/libcapi/libcapi.dll

lib /def:./capi/libcapi.def /name:libcapi.dll /out:./call_capi/libcapi/libcapi.lib /MACHINE:X64

sed -i "s/__SIZE_TYPE__/size_t/g" ./call_capi/libcapi/libcapi.h
sed -i "s/float\ _Complex/_Fcomplex/g" ./call_capi/libcapi/libcapi.h
sed -i "s/double\ _Complex/_Dcomplex/g" ./call_capi/libcapi/libcapi.h
sed -i "/GoComplex64/i #include\ <complex.h>" ./call_capi/libcapi/libcapi.h
