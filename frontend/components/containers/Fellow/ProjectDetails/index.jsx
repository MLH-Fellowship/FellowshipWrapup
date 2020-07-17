import fetch from "isomorphic-fetch";

import LogoTwitter from "react-ionicons/lib/LogoTwitter";
import LogoGithub from "react-ionicons/lib/LogoGithub";
import LogoLinkedin from "react-ionicons/lib/LogoLinkedin";
import IosMailOutline from "react-ionicons/lib/IosMailOutline";

function ProjectDetails({ accountInfo, contributions }) {
  const { Name, Location, Url, TwitterUsername } = accountInfo;
  const uniqueLangs = new Set();

  // Adding each individual language to our set so we can have unique values
  contributions.map((el) => uniqueLangs.add(el.PrimaryLanguage.Name));

  return (
    <div style={{ marginBottom: "11.5rem" }}>
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
                    from {Location}.
                  </>
                )}
              </h2>
              <p className="about-p mb-5">
                As part of the inaugural class of MLH Fellows, where I
                contributed to projects like:{" "}
                {contributions.map((el, i) => (
                  <span>
                    <a
                      href={`${el.Url}`}
                      style={{ textDecoration: "underline", fontWeight: "700" }}
                    >
                      {el.Name}
                    </a>
                    {i < contributions.length - 1 ? ", " : ""}
                  </span>
                ))}
                , and used languages like:{" "}
                {[...uniqueLangs].map((el, i) => (
                  <>
                    <span
                      style={{
                        color: "var(--color-green-light)",
                        fontWeight: "700",
                      }}
                    >
                      {el}
                    </span>{" "}
                    {i < [...uniqueLangs].length - 1 ? ", " : ""}
                  </>
                ))}
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
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ProjectDetails;
