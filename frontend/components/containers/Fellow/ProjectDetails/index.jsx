import fetch from "isomorphic-fetch";

import LogoTwitter from "react-ionicons/lib/LogoTwitter";
import LogoGithub from "react-ionicons/lib/LogoGithub";
import LogoLinkedin from "react-ionicons/lib/LogoLinkedin";
import IosMailOutline from "react-ionicons/lib/IosMailOutline";

function ProjectDetails({ accountInfo }) {
  const { Name, Location, Url, TwitterUsername } = accountInfo;

  return (
    <>
      <h4 className="display-4">A little about me</h4>
      <div className="row">
        <div className="col-lg-6 wrap text-left">
          <div className="about-desc">
            <h1 className="bold-text">About</h1>
            <div className="pt-5">
              <h2 className="mb-5">
                Hello, I'm {Name.split(" ")[0]}
                {Location && (
                  <>
                    <br />
                    from Zapopan, Jalisco.
                  </>
                )}
              </h2>
              <p className="about-p mb-5">
                As part of the inaugural class of MLH Fellows, I contributed to
                the [projectLanguage] ecosystem with a team of Fellows under the
                educational mentorship of professional software engineers.
              </p>
              <ul className="social-links list-unstyled mt-4">
                {Url && (
                  <li>
                    <a href={`${Url}`} target="_blank">
                      <LogoGithub></LogoGithub>
                    </a>
                  </li>
                )}

                {TwitterUsername && (
                  <li>
                    <a
                      href={`https://twitter.com/${TwitterUsername}`}
                      target="_blank"
                    >
                      <LogoTwitter></LogoTwitter>
                    </a>
                  </li>
                )}

                <li>
                  <a href="#">
                    <IosMailOutline></IosMailOutline>
                  </a>
                </li>
                <li>
                  <a href="#">
                    <LogoLinkedin></LogoLinkedin>
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default ProjectDetails;
