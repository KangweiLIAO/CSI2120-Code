/**
 * @author Kangwei Liao (8568800)
 */
package matching;

import java.io.IOException;
import java.util.ArrayList;

public interface TableWriter {
    public void writeTable(ArrayList<String> list, String fileName, int n) throws IOException;
}
