export default function ProgressTriangle({ issues, prs }) {
  const totalIssues = Object.keys(issues).length;
  const totalPrs = Object.keys(prs).length;
  // we'll assume that all issues closed were solved (change query from server)
  const solvedIssues = issues.reduce((acc, el) => {
    if (el.Closed) acc += 1;
    return acc;
  }, 0);
  const mergedPrs = prs.reduce((acc, el) => {
    if (el.Merged) acc += 1;
    return acc;
  }, 0);

  return (
    <div className="container-fluid">
      <h1 className="col-lg-8 bold-text mt-5 mb-5">Performance Stats</h1>
      <div className="col-md-12">
        <h3 className="progressTriangle-title">
          <span style={{ color: "#FF1140" }}>{mergedPrs}</span> PRs merged /{" "}
          <span style={{ color: "#FF1140" }}>{totalPrs}</span> total PRs
        </h3>
        <div className="progressTriangle">
          <div
            className="progressTriangle-bar"
            style={{ width: `${(mergedPrs * 100) / totalPrs}%` }}
          >
            <div className="progressTriangle-value">
              {((mergedPrs * 100) / totalPrs).toFixed(2)}%
            </div>
          </div>
        </div>

        <h3 className="progressTriangle-title">Merged/ Total Commits</h3>
        <div className="progressTriangle yellow">
          <div className="progressTriangle-bar" style={{ width: "87%" }}>
            <div className="progressTriangle-value">87%</div>
          </div>
        </div>

        <h3 className="progressTriangle-title">
          <span style={{ color: "#1E539F" }}>{solvedIssues}</span> solved issues
          / <span style={{ color: "#1E539F" }}>{totalIssues}</span> total issues
        </h3>
        <div className="progressTriangle blue">
          <div
            className="progressTriangle-bar"
            style={{ width: `${(solvedIssues * 100) / totalIssues}%` }}
          >
            <div className="progressTriangle-value">
              {((solvedIssues * 100) / totalIssues).toFixed(2)}%
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
