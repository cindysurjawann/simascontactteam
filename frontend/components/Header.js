import React from "react";
import Image from "next/image";
import Logo from "~/public/logo.png";
import { useRouter } from "next/router";

export default function Header() {
  const router = useRouter();

  const home = () => {
    router.push("/");
  };

  return (
    <div>
      <div style={{ paddingLeft: "20px", paddingTop: "10px" }}>
        <Image src={Logo} width="200px" height="50px" layout="fixed" onClick={home} alt="Logo Sinarmas" />
      </div>
      <div style={{ background: "#CC100F", width: "100%", height: "50px" }}>
        <div className="d-flex flex-row-reverse me-5 align-items-center " style={{ height: "100%" }}>
          <a href="#layananinfo" className="pe-2 text-white px-5" style={{ fontWeight: "bold", fontSize: "20px" }} onClick={home}>
            Pusat Informasi
          </a>
          <a href="#layanancs" className="pe-2 text-white px-5" style={{ fontWeight: "bold", fontSize: "20px" }} onClick={home}>
            Layanan CS
          </a>
        </div>
      </div>
    </div>
  );
}
