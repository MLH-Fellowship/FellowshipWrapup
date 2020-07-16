import MdHeart from "react-ionicons/lib/MdHeart";
import IosArrowUp from "react-ionicons/lib/IosArrowUp";
import IosCodeWorking from "react-ionicons/lib/IosCodeWorking";

export default function Footer() {
  return (
    <div className="container mt-5 mb-5">
    <hr className="p-1" />
        <div className="row slider-text align-items-center text-center justify-content-center"
            data-scrollax-parent="true">
            <div className="col-md-12">
                <h2>Hope you had a great time!</h2>
                <h1 className="display-3 mt-3 mb-3 footer-heading"><MdHeart fontSize="60px" color="red" beat={true} /> from 
                <img className="img-fluid ml-3" style={{width: '256px'}} src="/mlh-fellowship-banner.png" />
                </h1>
                <a className="display-1 mb-5" href="#top">
                    <IosArrowUp fontSize="60px"></IosArrowUp>
                </a>
            </div>
        </div>
        <div className="container-fluid text-right">
            <a href="https://github.com/MLH-Fellowship/FellowshipWrapup" className="footer-link"><IosCodeWorking fontSize="30px"></IosCodeWorking></a>
        </div>
    </div>
  );
}
