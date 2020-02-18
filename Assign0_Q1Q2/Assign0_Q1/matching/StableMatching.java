/**
 * @author Kangwei Liao (8568800)
 */
package matching;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.Queue;

public class StableMatching {
	
	private Table table1;
	private Table table2;
	private Table outputTable;
	
	/**
	 * Constructor
	 */
	
	public StableMatching(String table1Name, String table2Name) {
		table1 = new Table(table1Name);
		table2 = new Table(table2Name);
		outputTable = new Table();
	}
	
    /**
     * @throws IOException 
	 */
    
    private void galeShapley() throws IOException {
    	Queue<String> offer = new LinkedList<>();
    	HashMap<String, ArrayList<String>> employer = table1.readTableElem();
    	HashMap<String, ArrayList<String>> student = table2.readTableElem();
    	HashMap<String, String> pairs = new HashMap<String, String>();
    	int numOfPairs = table1.getRows() == table2.getRows() ? table1.getRows():-1;
    	
    	for (int i = 0; i < numOfPairs; i++) {
    		pairs.put(table2.getMainRoles()[i], null);
    		offer.add(table1.getMainRoles()[i]);
    	}
    	// obtains preference data from table1 and table2
		while (!offer.isEmpty()) {
			String currE = offer.poll();	// employer that has no student
			String currS = employer.get(currE).get(0);	// most preferred student
			for (int i = 0; i < numOfPairs; i++) {
				currS = employer.get(currE).get(i);	// loop student
				if (pairs.get(currS) == null) {
					// if currS has no job
					pairs.replace(currS, currE);
					break;
				} else if (student.get(currS).indexOf(pairs.get(currS)) > student.get(currS).indexOf(currE)) {
					// if currS prefer currE
					offer.add(pairs.get(currS));
					pairs.replace(currS, currE);
					break;
				}
			}
		}
		outputTable.configTable(pairs);
		String outputName = "matches_java_" + numOfPairs + "x" + numOfPairs + ".csv";
		outputTable.exportTable(outputName, 2);
		System.out.println("\"" + outputName + "\" created successfully!");
    }
    
	/**
	 * @param args
	 */
    
	public static void main(String[] args) {
		StableMatching gs = new StableMatching(args[0], args[1]);
		try {
			gs.galeShapley();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
