export default function ProgressRound({ countIss, countPr, contributions }) {
  console.log(contributions);

  return (
    <div className="container mb-5">
      <h1 className="col-lg-8 bold-text mt-5 mb-5">My Most Used Languages</h1>
      <div className="row">
        <div className="col-md-3 col-sm-6 pb-5">
          <div className="progress-round red">
            <span className="progress-left">
              <span className="progress-round-bar"></span>
            </span>
            <span className="progress-right">
              <span className="progress-round-bar"></span>
            </span>
            <div className="progress-round-value">
              {countIss > 100 ? 85 : countIss}%
            </div>
          </div>
          <h4 className="pt-4 text-center bold-text">
            {contributions[0].PrimaryLanguage.Name}
          </h4>
        </div>

        <div className="col-md-3 col-sm-6 pb-5">
          <div className="progress-round yellow">
            <span className="progress-left">
              <span className="progress-round-bar"></span>
            </span>
            <span className="progress-right">
              <span className="progress-round-bar"></span>
            </span>
            <div className="progress-round-value">
              {countPr > 100 ? 85 : countPr}%
            </div>
          </div>
          <h4 className="pt-4 text-center bold-text">
            {contributions[1].PrimaryLanguage.Name}
          </h4>
        </div>
      </div>
    </div>
  );
}
