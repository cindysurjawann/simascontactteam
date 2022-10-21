import style from "./sidebarn.module.scss";
import Image from "next/image";
import Logo from "../../public/logo.png";
import HomeIcon from "@mui/icons-material/Home";
import AccountBoxIcon from "@mui/icons-material/AccountBox";
import LogoutIcon from "@mui/icons-material/Logout";
import  {useRouter} from "next/router";
const Sidebar = ({ toggleActive }) => {
  const route = useRouter()
  const logout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    console.log(localStorage.getItem("token"));
    route.push("/loginForm")
  };
  return (
    <div className={style.sidebar}>
      <div className={style.top}>
        <Image src={Logo} alt="logo" />
      </div>
      <hr />
      <div className={style.center}>
        <ul>
          <p className={style.title}>MAIN MENU</p>
          <li onClick={() => toggleActive("halamanutama")}>
            <HomeIcon className={style.icon} />
            <span>Halaman Utama</span>
          </li>
          <li onClick={() => toggleActive("managecs")}>
            <AccountBoxIcon className={style.icon} />
            <span>Akun CS</span>
          </li>
          {/* <li>
            <a href="">
              <InfoIcon className={style.icon} />
              <span>Informasi</span>
            </a>
          </li> */}
          <li>
            <div href="#" onClick={logout}>
              <LogoutIcon className={style.icon} />
              <span>Keluar</span>
            </div>
          </li>
        </ul>
      </div>
    </div>
  );
};

export default Sidebar;