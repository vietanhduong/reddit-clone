import React, { useState, useEffect } from "react";
import Topic from "components/topic";
import NavBar from "components/navbar";
import { fetch } from "utils/api";

const Home = () => {
    const [topics, setTopics] = useState([]);
    useEffect(() => {
        fetch().then(data => setTopics(data.content)).catch(err => console.log(err));
    }, [])
    return (
        <div className="content">
            <NavBar />
            <hr />
            <h1>Topics</h1>
            <hr />
            {topics && topics.map(topic => (<Topic key={topic.id} topic={topic} />))}
        </div>
    );
};

export default Home;