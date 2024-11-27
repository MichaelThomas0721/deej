# shell.nix
{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  buildInputs = [
    pkgs.libappindicator
    pkgs.webkitgtk
    pkgs.pkg-config
    pkgs.libayatana-appindicator
  ];

  # Optional: Set environment variables if needed (e.g., for development purposes)
  shellHook = ''
    echo "Development environment for GTK3, AppIndicator, and WebKit2GTK is ready!"
  '';
}
