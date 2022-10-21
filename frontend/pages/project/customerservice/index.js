import Sidebar from "../../../components/sidebarcs/sidebarn";
import style from "./index.module.scss";
import HalamanUtama from "../../../components/halamanutamacs/halamanutama";
import ManageWA from "../../../components/managewa/managewa";
import ManageZoom from "../../../components/managezoom/managezoom";
import { useEffect, useState } from "react";
import {useRouter} from "next/router";
import ZoomHistory from "../../../components/zoomhistory/zoomhistory";

export default function Index() {
  const [loading, setLoading] = useState(true);
  const [showActive, setShowActive] = useState("halamanutama");
  const route = useRouter()
  const toggleActive = (key) => setShowActive((active) => (active === key ? "halamanutama" : key));

  useEffect(() => {
    const token = localStorage.getItem("token");
    let user = localStorage.getItem("user");
    if (token == null || user == null) {
      console.log("logout");
      route.push('/loginForm')
      return;
    }
    user = JSON.parse(user);
    if (user.role != 2) {
      if (user.role == 1) {
        console.log("redirect");
        route.back()
        return;
      }
      console.log("load", user);
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
              {showActive === "managezoom" && <ManageZoom />}
              {showActive === "managewa" && <ManageWA />}
              {showActive === "zoomhistory" && <ZoomHistory />}
            </>
          )}
        </div>
      </div>
    </div>
  );
}
