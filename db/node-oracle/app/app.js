import util from "node:util";
import oracledb from "oracledb";

async function run() {
	const conn = await oracledb.getConnection({
		user: "demo",
		password: "demopwd",
		connectString: "localhost/FREEPDB1"
	});

	oracledb.outFormat = oracledb.OUT_FORMAT_OBJECT;

	const event_id = await insertEvent(conn, {
		event_type: "demo.event",
		event_data: JSON.stringify({ id: 1, name: "demo event" }),
		created_by: "demo"
	});

	const events = await getEvents(conn, event_id);
	logData(events);

	await conn.close();
}

await run();

async function insertEvent(conn, data) {
	data = {
		...data,
		event_id: {
			dir: oracledb.BIND_OUT,
			type: oracledb.NUMBER
		}
	};

	const result = await conn.execute(
		`
		INSERT INTO DEMO_EVENT (EVENT_TYPE, EVENT_DATA, CREATED_BY)
     	VALUES (:event_type, :event_data, :created_by)
     	RETURNING EVENT_ID INTO :event_id
	 	`,
		data,
		{ autoCommit: true }
	);

	return result.outBinds.event_id[0];
}

async function getEvents(conn, event_id) {
	const result = await conn.execute(
		"SELECT * FROM DEMO_EVENT WHERE EVENT_ID = :event_id",
		{ event_id },
		{
			fetchInfo: {
				EVENT_DATA: { type: oracledb.STRING }
			}
		}
	);

	return result.rows.map((r) => ({
		...r,
		EVENT_DATA: JSON.parse(r.EVENT_DATA)
	}));
}

function logData(data) {
	console.log(util.inspect(data, { depth: null }) + "\n");
}
