import React from "react";
import { Link } from "react-router-dom";
import { me, create } from "utils/api";
import { browserHistory } from "utils/history";

const Create = () => {
    const [content, setContent] = React.useState("");

    React.useEffect(() => {
        me().then(data => {
            if (!data || Object.keys(data).length === 0) {
                browserHistory.push("/login");
                return;
            }
            if (!data.content.admin) {
                browserHistory.push("/");
            }
        })
    }, []);

    const onCreate = () => {
        create(content).then(data => {
            if (!data) {
                browserHistory.push("/login");
                return;
            }
            alert("create topic successful!");
        }).catch(err => {
            if (err.response) {
                const data = err.response.data;
                alert(`failed: ${data.message}`);
            }
        })
    }
    return (
        <div className="content">
            <textarea onChange={e => setContent(e.target.value)}></textarea>
            <input type="submit" value="Create" onClick={onCreate}></input>
            <br />
            <Link to={"/"}> {"<- "} Go home</Link>
        </div>
    );
}

export default Create;
