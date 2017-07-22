import React, {Component} from 'react';

class EventTable extends Component {
  render() {
    return (
      <table className="table">
        <thead>
          <tr>
            <th>
              <abbr title="uuid">UUID</abbr>
            </th>
            <th>
              <abbr title="Created At">CreatedAt</abbr>
            </th>
            <th>
              <abbr title="Finish At">Finish At</abbr>
            </th>
            <th>
              <abbr title="Time Interval">Time Interval</abbr>
            </th>
            <th>
              <abbr title="Driver Name">Driver Name</abbr>
            </th>
            <th>
              <abbr title="Driver device id">Device ID</abbr>
            </th>
            <th>
              <abbr title="Driver access token">Device access token</abbr>
            </th>
            <th>
              <abbr title="Action">Action</abbr>
            </th>
            <th>
              <abbr title="Status">Status</abbr>
            </th>
          </tr>
        </thead>
        <tbody>
          {this
            .props
            .events
            .map((item, idx) => {
              return <tr key={idx}>
                <th key={item.uuid}>{item.uuid}</th>
                <td key={item.createdAt}>{item.createdAt}</td>
                <td key={item.finishAt}>{item.finishAt}</td>
                <td key={item.timeInterval}>{item.timeInterval}</td>
                <td key={item.driver.uuid}>{item.driver.uuid}</td>
                <td key={item.driver.deviceId}>{item.driver.deviceId}</td>
                <td key={item.driver.accessToken}>{item.driver.accessToken}</td>
                <td key={item.action}>{item.action}</td>
                <td key={item.status}>{item.status}</td>
              </tr>
            })}
        </tbody>
      </table>
    )
  }
}

class UserList extends Component {
  render() {
    return (
      <select onChange={this.props.onChange}>
        <option>-- Select --</option>
        {this
          .props
          .users
          .map((item, idx) => {
            return <option value={item} key={idx}>{item}</option>
          })}
      </select>
    )
  }
}

class App extends Component {
  constructor(props) {
    super(props);
    this.addUser = this
      .addUser
      .bind(this);
    this.getEvents = this
      .getEvents
      .bind(this);
    this.state = {
      users: [],
      events: []
    }
  }
  getEvents(e) {
    var id = e.target.value;
    fetch("http://localhost:1323/user/" + id + "/events").then((res) => {
      return res.json();
    }).then((data) => {
      this.setState({events: data})
    })
  }
  getUsers() {
    fetch("http://localhost:1323/user").then((res) => {
      return res.json();
    }).then((data) => {
      this.setState({users: data})
    })
  }
  addUser() {
    fetch("http://localhost:1323/user", {method: "POST"}).then((res) => {
      return res.json();
    }).then((data) => {
      this.getUsers();
    })
  }
  componentDidMount() {
    this.getUsers();
  }
  render() {
    return (
      <div>
        <div className="columns">
          <div className="column">
            <div className="field has-addons">
              <p className="control is-expanded">
                <span className="select is-fullwidth">
                  <UserList onChange={this.getEvents} users={this.state.users}/>
                </span>
              </p>
              <p className="control">
                <button onClick={this.addUser} className="button">+</button>
              </p>
            </div>
          </div>
        </div>
        <div className="column">
          <EventTable events={this.state.events}/>
        </div>
      </div>
    );
  }
}

export default App;
