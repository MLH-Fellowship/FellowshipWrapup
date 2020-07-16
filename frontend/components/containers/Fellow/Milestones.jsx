import Countup from "react-countup";

export default function Milestones() {
  return (
    // Keeping the old code in until designs are finalized
    // <StyledBackground>
    //   <div className="text-container">
    //     <h1>Some of the milestones I hit during the fellowship</h1>
    //   </div>

    //   <div className="list">
    //     <div className="list-item">
    //       <span>
    //         <Countup end={30} duration={2} />
    //       </span>{" "}
    //       <br />
    //       lines of code
    //     </div>
    //     <div className="list-item">
    //       <span>
    //         <Countup end={30} duration={3} />
    //       </span>{" "}
    //       <br /> commits made
    //     </div>
    //     <div className="list-item">
    //       <span>
    //         <Countup end={30} duration={4} />
    //       </span>{" "}
    //       <br />
    //       issues participated
    //     </div>
    //   </div>
    // </StyledBackground>
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
              <Countup end={37} duration={5} />
            </strong>
          </div>
        </div>
      </div>
    </section>
  );
}
