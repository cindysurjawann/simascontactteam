import { useEffect, useState } from "react";
import styles from "../styles/Login.module.css";

import Router, {useRouter} from "next/router";
import Image from "next/image";
export default function LoginForm() {
  const [kode, setKode] = useState("")
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [timeLeft, setTimeLeft] = useState(null);
  const [isOTPSend, setIsOTPSend] = useState(false)
  const [load, setLoad] = useState(false)
  const route = useRouter()
  const current = new Date();
  const month = '0' + (current.getMonth() + 1);
  const day = '0' + current.getDate();
  const hours = '0' + current.getHours();
  const minutes = '0' + current.getMinutes();
  const second = '0' + current.getSeconds();
  const dateRequest = current.getFullYear() + '-'
              + month.substring(month.length-2,month.length) + '-'
              + day.substring(day.length-2,day.length) + 'T'
              + hours.substring(hours.length-2,hours.length) + ':'
              + minutes.substring(minutes.length-2,minutes.length) + ':'
              + second.substring(second.length-2,second.length) + '.000Z';

  useEffect(() => {
    let user = localStorage.getItem("user");
    let token = localStorage.getItem("token");
    if (user != null && token != null) {
      user = JSON.parse(user);
      if (user.role == 1) {
        Router.replace("/project/admin");
        return;
      } else if (user.role == 2) {
        Router.replace("/project/customerservice");
        return;
      }
    }
    localStorage.removeItem("user");
    localStorage.removeItem("token");
  }, []);
  useEffect(() => {
    // exit early when we reach 0
    if (!timeLeft) return;

    // save intervalId to clear the interval when the
    // component re-renders
    const intervalId = setInterval(() => {
      setTimeLeft(timeLeft - 1);
    }, 1000);

    // clear interval on re-render to avoid memory leaks
    return () => clearInterval(intervalId);
    // add timeLeft as a dependency to re-rerun the effect
    // when we update it
  }, [timeLeft]);
  
  async function sendOTP(){
    if(username == ""){
      alert("Please Masukkan username")
      return
    }
    try{
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}send-otp`, {
        method: "POST",
        header: {
        },
        body: JSON.stringify({"username":username}),
      });
      
      if(res.status != 200){
        throw "Gagal dapatin kode otp"()
      }
      setIsOTPSend(true)
      // responseMessage = await res.json()
      setTimeLeft(60*2)
      // alert("OTP telah dikirim, expired 5 menit")
    }catch(err){
      alert("Gagal mendapatkan OTP")
    }
    setLoad(false)
   
  }
  async function doLogin() {
    if(load) return
    setLoad(true)
    const body = {
      username: username,
      password: password,
      code: kode
    };
    try{
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}login`, {
        method: "POST",
        header: {
        },
        body: JSON.stringify(body),
      });
      const data = await res.json();
      if(res.status != 200){
        if(data.message == "OTP is wrong"){
          alert("Kode OTP is wrong")
          setLoad(false)
          return
        }else if(data.message == "OTP is expire"){
          alert("Kode OTP is expire")
          setLoad(false)
          return
        }
      }
      if (data.token) {
        localStorage.setItem("token", data.token);
  
        localStorage.setItem("user", JSON.stringify(data.data));
        postLastLogin();
        if (data.data.role == 1) {
          route.push("/project/admin")
        } else if (data.data.role == 2) {
          route.push("/project/customerservice")
        } else {
          console.log("Tidak ada Role");
        }
        
      } else {
        alert("Username / Password / OTP salah");
        setLoad(false)
      }
    
    }catch(err){
      alert("Server sedang bermasalah")
      setLoad(false)
    }
    
  }
  let minute,seconds = 0
  if(timeLeft){
    minute = Math.floor(timeLeft / 60).toLocaleString('en-US', {
      minimumIntegerDigits: 2,
      useGrouping: false
    })
    seconds = (timeLeft - minute * 60).toLocaleString('en-US', {
      minimumIntegerDigits: 2,
      useGrouping: false
    })
  }
  let css={
    cursorLogin:""
  }
  if(!isOTPSend){
    css.cursorLogin = "default"
  }else if(load){
    css.cursorLogin="wait"

  }else{
    css.cursorLogin=""
  }
  
  async function postLastLogin() {
    const data_user  = localStorage.getItem('user');
    const newData = JSON.parse(data_user);
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}updatelastlogin`, {
        method: "POST",
        headers: {
          Authorization: localStorage.getItem("token"),
        },
        body: JSON.stringify({
          username: newData.username,
          lastlogin: dateRequest,
        }),
      });
      if(res.status != 200){
        throw "gagal mendapatkan response update last login"();
      }
    } catch (error) {
      alert("Update Last Login Gagal");
    }
  }

  return (
    <div>
      <div className={styles.background}>
        <div className={styles.shape} />
        <div className={styles.shape} />
      </div>
      <div className="text-center mt-4">
        <Image src={"/logo.png"} width="378" height="100" alt="Logo" />

        <h2>Simas Contact dan Info</h2>
      </div>
      <div  id="formid" className={styles.form}>
        <h3>Masuk</h3>

        <label htmlFor="username" className={styles.label}>
          Username
        </label>
        <input type="text" placeholder="Masukan Username" id="username" value={username} onChange={(e)=>setUsername(e.target.value)} className={styles.input} />

        <label htmlFor="password">Password</label>
        <input type="password" placeholder="Password" id="password" value={password} onChange={(e)=>setPassword(e.target.value)} className={styles.input} />
        {
          timeLeft? (<label  >
            sisa : {minute } : {seconds}
          </label>):<></>
        }
        
        <div className="row">
          <div className="col-md-8">
          <input type="text" max="6" placeholder="Kode OTP" onChange={(e)=>setKode(e.target.value)} value={kode} className={styles.input} />
          </div>
          <div className="col-md-4">
            <button className="btn btn-success" disabled={timeLeft?true:false} onClick={sendOTP} style={{marginTop:"8px",height:"50px", width:"100%"}}>Send</button>
          </div>
        </div>
       
        {/* <a href="#" style={{ marginLeft: "70%", color: "#4A8CFF" }}>
          Lupa Kata Sandi ?
        </a> */}
        <button className={styles.button} onClick={doLogin} style={{cursor:css.cursorLogin}}  disabled={!isOTPSend} >Masuk</button>
      </div>
    </div>
    
  );
}
