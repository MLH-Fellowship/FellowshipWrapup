import fetch from "isomorphic-fetch";
import Countup from "react-countup";

function Milestones({ issues }) {
  const nIssues = Object.keys(issues).length;

  return (
    <section className="stats-section stats-counter mt-5">
      <div className="text-left stats-animate">
        <h1 className="display-1 section-heading">Milestones Hit</h1>
        <h1 className="bold-text bg-text">Achievements This Summer</h1>
      </div>
      <div className="row d-flex justify-content-start">
        <div className="col-md-7 col-sm-7">
          <div className="text">
            <span className="stats-label">lines of code</span>
            <strong className="number" id="countLOC">
              <Countup end={134} duration={5} />K
            </strong>
          </div>
        </div>
      </div>
      <div className="row d-flex justify-content-center">
        <div className="col-md-7 col-sm-7">
          <div className="text">
            <span className="stats-label">commits made</span>
            <strong className="number" id="countCommits">
              <Countup end={143} duration={5} />
            </strong>
          </div>
        </div>
      </div>
      <div className="row d-flex justify-content-end">
        <div className="col-md-7  col-sm-7">
          <div className="text">
            <span className="stats-label">issues opened</span>
            <strong className="number" id="countIssues">
              <Countup end={nIssues} duration={5} />
            </strong>
          </div>
        </div>
      </div>
    </section>
  );
}

Milestones.getInitialProps = async ({ query }) => {
  const res = await fetch(
    `${process.env.BACKEND_URL}/issuescreated/${query.uid}`,
    {
      method: "POST",
      body: JSON.stringify({
        secret: `${process.env.BACKEND_SECRET}`,
      }),
    }
  ).then((res) => res.json());

  return {
    info: res,
  };
};

export default Milestones;
