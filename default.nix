{pkgs, ...}:
pkgs.buildGoModule {
  pname = "linux-eva";
  version = "0.1.0";
  src = ./.;
  vendorHash = "sha256-5BUuzqd/rgPz+hddFz+x92I8sjd2TKKRR2QpW9ymXwA=";

  meta = {
    description = "Half bot half linux evangelist";
    homepage = "https://github.com/csl-club/linux-eva";
  };
}
