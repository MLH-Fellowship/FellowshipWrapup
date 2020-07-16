import { useRouter } from "next/router";

import Header from "../../components/containers/User/Header";

export default function User() {
  return (
    <div clasName="container">
      <Header />
      <section>
        <h4 className="display-4">A little about me,</h4>
        <div className="row">
          <div className="col-lg-6 wrap text-left">
            <div className="about-desc">
              <h1 className="bold-text">About</h1>
              <div className="pt-5">
                <h2 className="mb-5">Hello!</h2>
                <p>
                  I'm Sebastian from Zapopan. As part of the inaugural className
                  of MLH Fellows, I contributed to the JavaScript ecosystem with
                  a team of Fellows under the educational mentorship of
                  professional software engineers.
                </p>
                <p>
                  <a href="#">Add link to resume?</a>
                </p>
                <ul className="social-links list-unstyled mt-4">
                  <li>
                    <a href="#">
                      <ion-icon name="mail-outline"></ion-icon>
                    </a>
                  </li>
                  <li>
                    <a href="#">
                      <ion-icon name="logo-github"></ion-icon>
                    </a>
                  </li>
                  <li>
                    <a href="#">
                      <ion-icon name="logo-linkedin"></ion-icon>
                    </a>
                  </li>
                  <li>
                    <a href="#">
                      <ion-icon name="logo-twitter"></ion-icon>
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
          <div className="col-lg-6 wrap align-content-center">
            <div className="row mt-5 flex-column">
              <div className="col-md-8">
                <h2 className="mb-4">Most used languages</h2>
              </div>
              <div className="col-md-6 mt-4">
                <div className="progress-wrap">
                  <h4>JavaScript</h4>
                  <div className="progress progress-style">
                    <div
                      className="progress-bar color-1"
                      role="progressbar"
                      aria-valuenow="75"
                      aria-valuemin="0"
                      aria-valuemax="100"
                      style={{ width: "64%" }}
                    >
                      <span>Used in 78% commits</span>
                    </div>
                  </div>
                </div>
              </div>
              <div className="col-md-6 mt-4" data-animate-effect="fadeInRight">
                <div className="progress-wrap">
                  <h4>TypeScript</h4>
                  <div className="progress progress-style">
                    <div
                      className="progress-bar color-1"
                      role="progressbar"
                      aria-valuenow="60"
                      aria-valuemin="0"
                      aria-valuemax="100"
                      style={{ width: "64%" }}
                    >
                      <span>Used in 64% commits</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
      <section
        className="stats-section stats-counter mt-5"
        id="section-counter"
      >
        <div className="container">
          <div className="col-md-7 text-left heading-section stats-animate">
            <h1 className="display-1">Milestones Hit</h1>
            <h1 className="bold-text bg-text">Achievements This Summer</h1>
          </div>
          <div className="row d-flex justify-content-start">
            <div className="col-md-5 col-sm-5 counter-wrap stats-animate">
              <div className="text">
                <span className="stats-label">1000 lines of code</span>
                <strong className="number" id="countLOC">
                  0
                </strong>
              </div>
            </div>
          </div>
          <div className="row d-flex justify-content-center">
            <div className="col-md-5 col-sm-5 counter-wrap stats-animate">
              <div className="text">
                <span className="stats-label">commits made</span>
                <strong className="number" id="countCommits">
                  0
                </strong>
              </div>
            </div>
          </div>
          <div className="row d-flex justify-content-end">
            <div className="col-md-5 counter-wrap stats-animate">
              <div className="text">
                <span className="stats-label">issues participated in</span>
                <strong className="number" id="countIssues">
                  0
                </strong>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}
