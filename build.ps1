param($mode)

if ([String]::IsNullOrEmpty($mode)) {
    Write-Output "mode not set, using 's'"
    $mode = "s"
}

if ($mode -eq "a") {
    go build -buildmode=c-archive -ldflags "-s -w" -o ./call_capi/libcapi/libcapi.a ./capi
    gcc ./capi/libcapi.def ./call_capi/libcapi/libcapi.a -shared -lwinmm -lWs2_32 -o ./call_capi/libcapi/libcapi.dll
}
elseif ($mode -eq "s") {
    go build -buildmode=c-shared -ldflags "-s -w" -o ./call_capi/libcapi/libcapi.dll ./capi
}

lib /def:./capi/libcapi.def /name:libcapi.dll /out:./call_capi/libcapi/libcapi.lib /MACHINE:X64

sed -i "s/__SIZE_TYPE__/size_t/g" ./call_capi/libcapi/libcapi.h
sed -i "s/float\ _Complex/_Fcomplex/g" ./call_capi/libcapi/libcapi.h
sed -i "s/double\ _Complex/_Dcomplex/g" ./call_capi/libcapi/libcapi.h
sed -i "/_Fcomplex\ GoComplex64/i #include\ <complex.h>" ./call_capi/libcapi/libcapi.h
