import React from "react";

// reactstrap components
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import { Modal, ModalBody } from "reactstrap";
import style from "./berhasilsubmit.module.scss";

function BerhasilSubmit(props) {
  return (
    <>
      <Modal toggle={props.close} isOpen={props.show}>
        <ModalBody className={style.ModalBody}>
          <div className={style.content}>
            <CheckCircleIcon className={style.icon} />
            <h3>Data Berhasil Diubah.</h3>
          </div>
        </ModalBody>
      </Modal>
    </>
  );
}

export default BerhasilSubmit;
