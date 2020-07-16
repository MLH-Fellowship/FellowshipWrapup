export default function ProgressTriangle() {
    return (
        <div className="container-fluid">
            <h1 className="col-lg-8 bold-text mb-5">PR's merged</h1>
            <div className="col-md-12">
                <h3 className="progressTriangle-title">Merged/ Total PRs</h3>
                <div className="progressTriangle">
                    <div className="progressTriangle-bar" style={{width:'75%'}}>
                        <div className="progressTriangle-value">75%</div>
                    </div>
                </div>

                <h3 className="progressTriangle-title">Merged/ Total Commits</h3>
                <div className="progressTriangle yellow">
                    <div className="progressTriangle-bar" style={{width:'87%'}}>
                        <div className="progressTriangle-value">87%</div>
                    </div>
                </div>

                <h3 className="progressTriangle-title">Resolved/ Total Issues</h3>
                <div className="progressTriangle blue">
                    <div className="progressTriangle-bar" style={{width:'57%'}}>
                        <div className="progressTriangle-value">57%</div>
                    </div>
                </div>
            </div>
        </div>
    )
}