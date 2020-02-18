/**
 * @author Kangwei Liao (8568800)
 */
package matching;

import java.util.ArrayList;

public interface TableReader {
	
    public ArrayList<String> readTable() throws Exception;
    
	public int getRows();
}
