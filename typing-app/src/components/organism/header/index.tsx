import Image from "next/image";
import "../../../styles/header.scss";

const Header = () => {
  return (
    <div className="header">
      <div className="logo">
        <img src="/img/header_logo.png" />
      </div>
      <div className="user">
        <div className="icon">
          <img src="/img/icon_default.png" />
        </div>
        <div className="name">名称未設定</div>
        <div className="number">
          <div className="tip">Student No.</div>
          <div className="value">00000000</div>
        </div>
      </div>
    </div>
  );
};

export default Header;
