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
          <div className="text">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus vel massa pulvinar, semper arcu porttitor,
            sodales dui. Nam vitae blandit quam. Sed condimentum euismod placerat. Fusce id ipsum ante. Praesent
            pulvinar, urna at tempor pellentesque, erat ligula lobortis metus, ut ultrices ipsum nunc non turpis. Nunc
            egestas urna ipsum, dignissim porta orci rutrum sed. Etiam in tristique urna. Fusce eu eros laoreet, varius
            ipsum in, eleifend dui. Proin dapibus tortor nec ultricies porta. Suspendisse potenti. Suspendisse potenti.
            Donec vel volutpat arcu. Morbi ullamcorper a velit finibus placerat. Ut ac metus vitae lectus ornare
            fermentum vitae vitae sem. Morbi laoreet finibus purus nec faucibus.{" "}
          </div>
        </div>
      </div>
      <Footer></Footer>
    </div>
  );
}
