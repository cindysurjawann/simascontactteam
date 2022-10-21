import style from "./halamanutama.module.scss";
import Image from "next/image";
import foto1 from "../../public/assets/info1.jpg";
import foto2 from "../../public/assets/info2.jpg";

import {useEffect, useState} from 'react'


const HalamanUtama = () => {
  const [loc, setLoc] = useState(null)
  const [date, setDate] = useState(null)
  const data_user  = localStorage.getItem('user')
  const newData = JSON.parse(data_user)
  useEffect(() => {
    const data_user  = localStorage.getItem('user')
    const newData = JSON.parse(data_user)
    if(newData != null && newData != undefined){
      if(newData["lastlogin"] != null && newData["lastlogin"] != undefined){
        const dateRes = newData.lastlogin.substring(0,10)
        const timeRes = newData.lastlogin.substring(11,19)
        const date = dateRes + " " + timeRes
        setDate(date)
      }
     
    }
    const item = localStorage.getItem('location')
    const obj = JSON.parse(item);
    if(obj != undefined && obj != null){
      setLoc(obj)
    }
    if(localStorage.getItem('location') === null) {
      navigator.geolocation.getCurrentPosition(function(position) {
        getLocation(position.coords.latitude, position.coords.longitude)
      });
    }
  }, []);

  const getLocation = async (lat, lang) => {
    try {
      const newUrl = `https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=${lat}&longitude=${lang}&localityLanguage=en`
      const res = await fetch(newUrl);
      const data = await res.json();
      localStorage.setItem("location", JSON.stringify(data))
      setLoc(data)
    }
    catch (error) {
    }
  }

  return (
 
      <div className={style.utama}>
        <h1 className={style.title}>Selamat Datang {newData.name}</h1>
        <hr />
        {loc !== null ? 
          <div className={style.alert}>
          <p>Terakhir Login : {date}</p> 
          <p>Lokasi Sekarang : {loc.locality} {loc.city} {loc.principalSubdivision} {loc.countryName} </p> 
          </div> : ''
        }
        <br />
        <br />
        <div className={style.informasi}>
          <h3 style={{ fontSize: "24px", fontWeight: "450" }}>Informasi</h3>

          <Image src={foto1} width={700} height={325} alt="foto1" />

          <p style={{ fontSize: "20px", textAlign: "justify" }}>Waspada Penipuan, Begini Tips Transaksi Aman di ATM Bank Sinarmas</p>
          <br />
          <br />

          <Image src={foto2} width={700} height={325} alt="foto2" />
          <p style={{ fontSize: "20px", textAlign: "justify" }}>Amankan Kartu Kredit dengan Cara Freeze Lewat Aplikasi Simobi+</p>
        </div>
      </div>
    
  );
};

export default HalamanUtama;
