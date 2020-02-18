/**
 * @author Kangwei Liao (8568800)
 */
package matching;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;

public class Table {
	
	private ArrayList<String> mainRoles;
	private ArrayList<String> elements;
	private HashMap<String, ArrayList<String>> tableHashMap;
	private TableFileManager tManager;

	/**
	 * Constructors
	 */
	public Table() {
		mainRoles = new ArrayList<String>();
		tableHashMap = new HashMap<String, ArrayList<String>>();
		tManager = new TableFileManager();
		elements = new ArrayList<String>();
	}
	
	public Table(String tableName) {
		try {
			mainRoles = new ArrayList<String>();
			tableHashMap = new HashMap<String, ArrayList<String>>();
			tManager = new TableFileManager(tableName);
			elements = tManager.readTable();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	public HashMap<String, ArrayList<String>> readTableElem() {
		for (int i = 0; i < elements.size(); i++) {
			if(i%(tManager.getRows()+1) == 0) {
				ArrayList<String> sBuffer = new ArrayList<String>();
				for (int j = i; j < i+tManager.getRows(); j++) {
					sBuffer.add(elements.get(j+1));
				}
				tableHashMap.put(elements.get(i), sBuffer);
				mainRoles.add(elements.get(i));
			}
		}
		return tableHashMap;
	}
	
	public void configTable(HashMap<String, String> map) {
		for (String key : map.keySet()) {
			ArrayList<String> arr = new ArrayList<String>();
			arr.add(map.get(key));
			tableHashMap.put(key, arr);
		}
	}
	
    public void exportTable(String fileName, int gap) throws IOException {
    	ArrayList<String> output = new ArrayList<String>();
    	for (String pair : tableHashMap.keySet()) {
			output.add(tableHashMap.get(pair).get(0));
			output.add(pair);
		}
    	tManager.writeTable(output, fileName, gap);
    }
    
	public String[] getMainRoles() {
		String[] sBuffer = new String[mainRoles.size()];
		for (int i = 0; i<mainRoles.size(); i++) {
			sBuffer[i] = mainRoles.get(i);
		}
		return sBuffer;
	}
	
	public int getRows() {return tManager.getRows();}
}
