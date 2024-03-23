import "../../styles/main.scss";
import "../../styles/title.scss";

export default function Title() {
  return (
    <div className="title">
      <div className="header">
        <img src="/img/sukashi.png" className="sukashi"></img>
        <img src="/img/header_logo.png" className="logo"></img>
      </div>
      <div className="footer">
        <div className="button">画面をクリックしてスタート</div>
        <div className="information">
          <p>
            (c) 2024 Faculty of Informatics, Shizuoka University all rights reserved.
            <br />
            Developed by IT Solution Room, Shizuoka University
            <br />
            Frontend Version 1.0.0 / Backend Version 1.0.0
          </p>
        </div>
      </div>
    </div>
  );
}
