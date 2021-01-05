import React, { useState, useEffect } from "react";
import {vote} from "utils/api";

const Topic = ({ topic }) => {

    const [counter, setCounter] = useState(0)
    const upVote = () => {
        onVote(true);
    }

    const downVote = () => {
        onVote(false);
    }


    const onVote = (isUp) => {
        vote(topic.id, isUp).then(data => {
            const content = data.content;
            setCounter(content.up_vote-content.down_vote);
        }).catch(err => {
            if (err.response) {
                const data = err.response.data;
                if (data.code === 401) {
                    alert("please login again");
                }
            }
        })
    }
    useEffect(() => {
        setCounter(topic.up_vote - topic.down_vote);
    }, [topic])

    return (
        <div className="topic">
            <div className="left">
                <div className="arrow noselect" onClick={upVote}>
                    ∆
               </div>
                <div className="number">
                    {counter}
                </div>
                <div className="arrow noselect" onClick={downVote}>
                    ∇
               </div>
            </div>
            <div className="right">{topic.content}</div>
        </div>
    );
};


export default Topic;