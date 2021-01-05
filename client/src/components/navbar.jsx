import React, { useEffect } from "react";
import { Link } from "react-router-dom";
import { me } from "utils/api";

const Navbar = () => {
    const [profile, setProfile] = React.useState({});
    useEffect(() => {
        me().then(data => {
            setProfile(data.content);
        }).catch(err => {
            if (err.response) {
                const data = err.response.data;
                if (data.code === 401) {
                    localStorage.removeItem("access_token");
                }
            }
        })
    }, []);
    return (
        <ul className="nav">
            <li><Link to="/">Topics</Link></li>
            {Object.keys(profile).length !== 0 && (<>
                {profile.admin && <li><Link to="/create">Create</Link></li>}
                <li>{profile.full_name} ({profile.username})</li>
            </>)}
            {Object.keys(profile).length === 0 && (<>
                <li><Link to="/login">Login</Link></li>
            </>)}
        </ul>
    );
};



export default Navbar;