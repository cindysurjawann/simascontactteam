import Header from "../../components/Header";
import style from "./kontakresmi.module.scss";
import Image from "next/future/image";
import logofacebook from "../../public/logofacebook.png";
import logoinstagram from "../../public/logoinstagram.png";
import logotiktok from "../../public/logotiktok.png";
import logotwitter from "../../public/logotwitter.png";
import logoyoutube from "../../public/logoyoutbe.png";

const KontakResmi = () => {
  return (
    <div>
      <Header />
      <div>
        <h2 className="ms-3 my-4">Kontak Resmi</h2>
        <div className={style.container}>
          <table className="mt-5">
            <tr>
              <td>
                <a href="https://id-id.facebook.com/BankSinarmas/" target="_blank" rel="noopener noreferrer">
                  {" "}
                  <Image src={logofacebook} width={76} height={71} alt={"logo instagram"} />
                  @BankSinarmas
                </a>
              </td>
            </tr>
            <tr>
              <td>
                <a href="https://www.instagram.com/banksinarmas/" target="_blank" rel="noopener noreferrer">
                  {" "}
                  <Image src={logoinstagram} width={76} height={71} alt={"logo instagram"} />
                  @banksinarmas
                </a>
              </td>
            </tr>
            <tr>
              <td>
                <a href="https://twitter.com/BankSinarmas" target="_blank" rel="noopener noreferrer">
                  {" "}
                  <Image src={logotwitter} width={76} height={71} alt={"logo tiktok"} />
                  @BankSinarmas
                </a>
              </td>
            </tr>
            <tr>
              <td>
                <a href="https://www.tiktok.com/@banksinarmasofficial" target="_blank" rel="noopener noreferrer">
                  {" "}
                  <Image src={logotiktok} width={76} height={71} alt={"logo tiktok"} />
                  @banksinarmasofficial
                </a>
              </td>
            </tr>
            <tr>
              <td>
                <a href="https://www.youtube.com/c/BankSinarmasOfficial" target="_blank" rel="noopener noreferrer">
                  {" "}
                  <Image src={logoyoutube} width={76} height={71} alt={"logo youtube"} />
                  @BankSinarmas
                </a>
              </td>
            </tr>
          </table>
        </div>
      </div>
      <div>
        <div className={style.footer}>
          <p style={{}}>© 2022 Simas Contact dan Info</p>
        </div>
      </div>
    </div>
  );
};

export default KontakResmi;
