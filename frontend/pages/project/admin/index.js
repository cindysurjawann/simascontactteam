import Sidebar from "../../../components/sidebaradmin/sidebarn";
import style from "./index.module.scss";
import HalamanUtama from "../../../components/halamanutamaadmin/halamanutama";
import ManageCS from "../../../components/managecs/managecs";

import { useEffect, useState } from "react";
import {useRouter} from "next/router";

export default function Index() {
  const [loading, setLoading] = useState(true);
  const [showActive, setShowActive] = useState("halamanutama");
  const route = useRouter()
  const toggleActive = (key) =>
    setShowActive((active) => (active === key ? "halamanutama" : key));

  useEffect(() => {
    const token = localStorage.getItem("token");
    let user = localStorage.getItem("user");
    if (token == null || user == null) {
      console.log("logout");
      route.push('/loginForm')
      return;
    }
    user = JSON.parse(user);
    if (user.role != 1) {
      if (user.role == 2) {
        console.log("redirect");
        route.back()
        return;
      }
      route.push('/loginForm')
      return;
    }
    setLoading(false);
    console.log(token, user);
  }, []);

  return (
    <div className={style.home}>
      <Sidebar toggleActive={toggleActive} />
      <div className={style.homeContainer}>
        <div className={style.content}>
          {loading ? (
            <div>
              <h1>Please wait</h1>
            </div>
          ) : (
            <>
              {showActive === "halamanutama" && <HalamanUtama />}
              {showActive === "managecs" && <ManageCS />}
            </>
          )}
        </div>
      </div>
    </div>
  );
}