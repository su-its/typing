import "../../../styles/main.scss";
import "../../../styles/game.scss";
import Header from "../../../components/organism/header/";
import Footer from "../../../components/organism/footer/";

export default function Home() {
  return (
    <div>
      <Header></Header>
      <div className="main">
        <div className="box">
          <div className="heading" id="heading_name">
            Article Name
          </div>
          <div className="heading" id="heading_time">
            Time Remain
          </div>
          <div className="heading" id="heading_position">
            Progress
          </div>
          <div className="heading" id="heading_speed">
            Speed
          </div>
          <div className="progress" id="progress_time"></div>
          <div className="progress" id="progress_position"></div>
          <div className="progress" id="progress_speed"></div>
          <img id="gauge_time" src="/img/gauge_time.png" />
          <img id="gauge_position" src="/img/gauge_position.png" />
          <img id="gauge_speed" src="/img/gauge_speed.png" />
          <div className="title">Lorem Ipsum</div>
          <div className="text">
            <div>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus vel massa pulvinar, semper arcu
              porttitor, sodales dui. Nam vitae blandit quam. Sed condimentum euismod placerat. Fusce id ipsum ante.
              Praesent pulvinar, urna at tempor pellentesque, erat ligula lobortis metus, ut ultrices ipsum nunc non
              turpis. Nunc egestas urna ipsum, dignissim porta orci rutrum sed. Etiam in tristique urna. Fusce eu eros
              laoreet, varius ipsum in, eleifend dui. Proin dapibus tortor nec ultricies porta. Suspendisse potenti.
              Suspendisse potenti. Donec vel volutpat arcu. Morbi ullamcorper a velit finibus placerat. Ut ac metus
              vitae lectus ornare fermentum vitae vitae sem. Morbi laoreet finibus purus nec faucibus.
            </div>
          </div>
          <div id="info_time">
            残り <span>250</span> 秒
          </div>
          <div id="info_text">123 語 / 4567 字</div>
        </div>
      </div>
      <Footer></Footer>
    </div>
  );
}
