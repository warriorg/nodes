package me.warriorg.sbm.department.domain;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Date;
import lombok.Getter;
import lombok.Setter;

/**
* generated by Generate dcits
* 
* @author: 朱正东
* @date: 2019-02-11
*/
@Setter @Getter
@Table(name = "DEPARTMENTS")
public class Department implements Serializable {
  private static final long serialVersionUID = 1L;

  /**
  * null
  */
  @Column(name = "dept_no")
  private String deptNo;
  /**
  * null
  */
  @Column(name = "dept_name")
  private String deptName;
}