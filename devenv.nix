{
  pkgs,
  lib,
  config,
  inputs,
  ...
}:

{
  packages = [
    pkgs.git
    pkgs.go

    pkgs.pkg-config

    pkgs.xorg.libX11
    pkgs.xorg.libXrandr
    pkgs.xorg.libXinerama
    pkgs.xorg.libXcursor
    pkgs.xorg.libXi
    pkgs.xorg.libXfixes
    pkgs.xorg.libxcb

    pkgs.libxkbcommon

    pkgs.wayland

    pkgs.libGL
    pkgs.vulkan-headers
    pkgs.vulkan-loader
  ];

  languages.go.enable = true;

  env.LD_LIBRARY_PATH = lib.makeLibraryPath [
    pkgs.libGL
    pkgs.vulkan-loader
    pkgs.libxkbcommon

    pkgs.wayland

    pkgs.xorg.libX11
    pkgs.xorg.libXrandr
    pkgs.xorg.libXinerama
    pkgs.xorg.libXcursor
    pkgs.xorg.libXi
    pkgs.xorg.libxcb
  ];
}
