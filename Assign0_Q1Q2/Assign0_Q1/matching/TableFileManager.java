/**
 * @author Kangwei Liao (8568800)
 */
package matching;

import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.io.BufferedReader;
import java.io.BufferedWriter;

public class TableFileManager implements TableReader, TableWriter {
	
	private int numOfRows;
	private BufferedReader bfReader;
	private BufferedWriter bfWriter;
	
	/**
	 * Constructors
	 */

	public TableFileManager() {
		numOfRows = 0;
	}
	
	public TableFileManager(String fileName) {
		numOfRows = 0;
		try {
			bfReader = new BufferedReader(new FileReader(System.getProperty("user.dir") + "/" + fileName));
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	/**
	 * @throws Exception 
	 * 
	 * @return ArrayList<String>
	 */
    
    public ArrayList<String> readTable() throws Exception {
        String str;
    	ArrayList<String> arr = new ArrayList<String>();
		while ((str = bfReader.readLine()) != null) {
        	for (String s : str.split(",")) {
        		arr.add(s);
        	}
        	numOfRows++;
        }
        return arr;
    }
   
    /**
	 * @param ArrayList<String> list
	 * @param String fileName
	 * @param int n
	 * 
	 * @throws IOException 
	 */
    
    public void writeTable(ArrayList<String> list, String fileName, int n) throws IOException {
    	File file = new File(fileName);
		bfWriter = new BufferedWriter(new FileWriter(file, true));
    	for (int i = 0; i < list.size(); i++) {
    		if (i != 0 && i % n == 0) bfWriter.write("\n" + list.get(i) + ",");
    		else bfWriter.write(list.get(i) + ",");
    	}
    	bfWriter.close();
    }
    
	/**
	 * @return int
	 */
    
	public int getRows() {return numOfRows;}
}
