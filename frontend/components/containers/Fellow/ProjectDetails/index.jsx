// import {LogoTwitter, LogoGithub, LogoLinkedin, IosMailOutline} from "react-ionicons"; Error: Not found, I'm probably doing something really dumb

export default function ProjectDetails() {
  return (
    <>
     {/* <StyledBackground>
       <div className="text-container">
         <h1>
           During the past 3 months I've contributed to projects like{" "}
           <span>[project]</span> &<span>[project]</span>
         </h1>
       </div>
     </StyledBackground> */}
    <h4 className="display-4">A little about me</h4>
    <div className="row">
        <div className="col-lg-6 wrap text-left">
            <div className="about-desc">
                <h1 className="bold-text">About</h1>
                <div className="pt-5">
                    <h2 className="mb-5">Hello, I'm Sebastian<br/>from Zapopan, Jalisco.</h2> 
                    <p className="about-p mb-5">As part of the inaugural class of MLH Fellows, I
                    contributed to the [projectLanguage] ecosystem with a team of Fellows under the
                    educational mentorship of professional software engineers.
                    </p>
                    <ul className="social-links list-unstyled mt-4">
                        <li><a href="#">
                                {/* <IosMailOutline></IosMailOutline> */}
                            </a></li>
                        <li><a href="#">
                                {/* <LogoGithub></LogoGithub> */}
                            </a></li>
                        <li><a href="#">
                                {/* <LogoLinkedin></LogoLinkedin> */}
                            </a></li>
                        <li><a href="#">
                                {/* <LogoTwitter></LogoTwitter> */}
                            </a></li>
                    </ul>
                </div>
            </div>
        </div>
        {/* Removed until we figure out what to stuff here */}
        {/* <div className="col-lg-6 wrap align-content-center">
            <div className="row mt-5 flex-column">
                <div className="col-md-8">
                    <h2 className="mb-4">Most used languages in my projects</h2>
                </div>
                <div className="col-md-6 mt-4">
                    <div className="progress-wrap">
                        <h4>[projectLanguage1]</h4>
                        <div className="progress progress-style">
                            <div className="progress-bar color-1" role="progressbar" aria-valuenow="75"
                                aria-valuemin="0" aria-valuemax="100" style={{width:0.66}}>
                                <span>Used in 2/3rd projects</span>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="col-md-6 mt-4">
                    <div className="progress-wrap">
                        <h4>[projectLanguage2]</h4>
                        <div className="progress progress-style">
                            <div className="progress-bar color-1" role="progressbar" aria-valuenow="60"
                                aria-valuemin="0" aria-valuemax="100" style={{width:0.33}}>
                                <span>Used in 1/3rd projects</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div> */}
    </div>
    </>
  );
}
