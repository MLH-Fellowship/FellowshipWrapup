export default function ProgressTriangle({ issues }) {
  const totalIssues = Object.keys(issues).length;
  // we'll assume that all issues closed were solved (change query from server)
  const solvedIssues = issues.reduce((acc, el) => {
    if (el.Closed === true) acc += 1;
    return acc;
  }, 0);

  return (
    <div className="container-fluid">
      <h1 className="col-lg-8 bold-text mt-5 mb-5">PR's merged</h1>
      <div className="col-md-12">
        <h3 className="progressTriangle-title">Merged/ Total PRs</h3>
        <div className="progressTriangle">
          <div className="progressTriangle-bar" style={{ width: "75%" }}>
            <div className="progressTriangle-value">75%</div>
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
            style={{ width: `${(solvedIssues * 100) / totalIssues}` }}
          >
            <div className="progressTriangle-value">
              {(solvedIssues * 100) / totalIssues}%
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
